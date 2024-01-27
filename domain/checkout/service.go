package checkout

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pevin/internal-pos-service-api/domain/checkout/entity"
	"github.com/pevin/internal-pos-service-api/domain/checkout/rest"
	entityEmployee "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	entityUser "github.com/pevin/internal-pos-service-api/domain/user/entity"
	"github.com/rs/zerolog/log"
)

type Service struct {
	checkoutRepo checkoutRepo
	employeeRepo employeeRepo
	authService  authService
}

type ServiceOpt struct {
	CheckoutRepo checkoutRepo
	EmployeeRepo employeeRepo
	AuthService  authService
}

type checkoutRepo interface {
	Transact(co entity.Checkout, bal entityEmployee.Balance, newBal float64) error
}

type employeeRepo interface {
	GetEmployeeNumberFromRFID(rfid, companyID string) (string, error)
	GetEmployeeAndBalance(employeeNumber, companyID string) (entityEmployee.Employee, entityEmployee.Balance, error)
}

type authService interface {
	FromRequestContext(req events.APIGatewayProxyRequestContext) (entityUser.User, error)
}

func NewService(opt ServiceOpt) *Service {
	return &Service{
		checkoutRepo: opt.CheckoutRepo,
		employeeRepo: opt.EmployeeRepo,
		authService:  opt.AuthService,
	}
}

func (s *Service) Create(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	// Get user from request context
	u, err := s.authService.FromRequestContext(req.RequestContext)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user from request context.")
		return
	}

	// Get create request payload
	var rp rest.CreateRequestPayload
	err = json.Unmarshal([]byte(req.Body), &rp)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal create checkout request payload.")
	}

	// get employee number
	employeeNumber, err := s.employeeRepo.GetEmployeeNumberFromRFID(rp.RFID, u.CompanyID)
	if err != nil {
		log.Error().Err(err).Msg("Got error calling dynamodb Query.")
		return
	}

	if len(employeeNumber) == 0 {
		// todo: return bad request - RFID not registered.
		// res = rest.BadRequestResponse("Employee not registered.")
		res.StatusCode = 400
		res.Body = "RFID doesn't exist."
		return
	}

	// Get employee entity and balance entity from dynamodb
	emp, bal, err := s.employeeRepo.GetEmployeeAndBalance(employeeNumber, u.CompanyID)
	if err != nil {
		log.Error().Err(err).Msg("Got error when getting employee and balance entity.")
		return
	}

	if len(emp.EmployeeNumber) == 0 || len(bal.EmployeeNumber) == 0 {
		log.Warn().Msg("Invalid employee and balance entity.")
		// Employee isn't configured correctly - return bad request
		// todo: return bad request
		res.StatusCode = 400
		res.Body = "employee isn't configured correctly."
		return
	}

	// Init checkout
	now := time.Now()
	co := entity.Checkout{
		ID:               now.Format("20060102150405"),
		CompanyID:        u.CompanyID,
		ConcessionaireID: u.ConcessionaireID,
		EmployeeNumber:   employeeNumber,
		Station:          rp.Station,
		CreatedBy:        u.ToPartialUser(),
		CreatedAt:        now,
		DiscountAmount:   0, // todo: get discount from employee discount entity
	}

	items := make([]entity.CheckoutItem, len(rp.Items))
	for i, item := range rp.Items {
		items[i] = entity.CheckoutItem{
			MealCode:         item.Code,
			MealName:         item.Name,
			Price:            item.Price,
			Calories:         item.Calories,
			Quantity:         item.Quantity,
			Category:         item.Category,
			SubTotalPrice:    item.Price * float64(item.Quantity),
			SubTotalCalories: item.Calories * float64(item.Quantity),
		}
		co.TotalCalories += items[i].SubTotalCalories
		co.TotalGrossAmount += items[i].SubTotalPrice
	}

	co.CheckoutItems = items
	co.TotalNetAmount = co.TotalGrossAmount - co.DiscountAmount

	// Validate if employee has sufficient balance
	if co.TotalNetAmount > bal.Balance {
		// balance is not sufficient - return bad request
		// todo: return bad request
		res.StatusCode = 400
		res.Body = fmt.Sprintf("insufficient balance %f %f", co.TotalNetAmount, bal.Balance)
		return
	}

	// Transact checkout
	bal.UpdatedAt = now
	balAfter := bal.Balance - co.TotalNetAmount

	err = s.checkoutRepo.Transact(co, bal, balAfter)
	if err != nil {
		log.Error().Err(err).Msg("Got error when saving checkout and balance in")
	}

	// Prepare response payload
	resp := rest.CheckoutResponsePayload{
		Checkout: co,
		Employee: rest.EmployeeResponsePayload{
			EmployeeNumber: emp.EmployeeNumber,
			Name:           emp.GetEmployeeName(),
			ImageUrl:       emp.ImageUrl,
			BalanceBefore:  bal.Balance,
			BalanceAfter:   balAfter,
		},
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling checkout response payload.")
		return
	}

	res = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(respByte),
	}

	return
}
