package checkout_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pevin/internal-pos-service-api/domain/checkout"
	entityCheckout "github.com/pevin/internal-pos-service-api/domain/checkout/entity"
	"github.com/pevin/internal-pos-service-api/domain/checkout/mocks"
	restCheckout "github.com/pevin/internal-pos-service-api/domain/checkout/rest"
	entityEmployee "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	userEntity "github.com/pevin/internal-pos-service-api/domain/user/entity"
	"github.com/pevin/internal-pos-service-api/lib/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServiceCreateSuccess(t *testing.T) {
	checkoutRepo := new(mocks.CheckoutRepo)
	employeeRepo := new(mocks.EmployeeRepo)
	opts := checkout.ServiceOpt{
		CheckoutRepo: checkoutRepo,
		EmployeeRepo: employeeRepo,
	}
	service := checkout.NewService(opts)

	ctx := context.TODO()

	rfid := "12345"
	body := map[string]interface{}{
		"rfid":    rfid,
		"station": "test station",
		"items": []map[string]interface{}{
			{
				"name":     "Coffee",
				"code":     "COF",
				"category": "Drinks",
				"price":    100,
				"calories": 50,
				"qty":      2,
			},
		},
	}
	bodyStr, marshalErr := json.Marshal(body)
	require.NoError(t, marshalErr)

	// set mock response
	u := userEntity.User{
		Username:   "test-user",
		GivenName:  "john",
		FamilyName: "wick",
		CompanyID:  "company-id-1234",
	}
	req := rest.Request{
		Body: string(bodyStr),
		User: u,
	}

	employeeNumber := "test-employee-number"
	bal := entityEmployee.Balance{
		CompanyID:      u.CompanyID,
		EmployeeNumber: employeeNumber,
		Balance:        5000,
	}
	emp := entityEmployee.Employee{
		CompanyID:      u.CompanyID,
		EmployeeNumber: employeeNumber,
		FirstName:      "john",
		LastName:       "wick",
	}

	employeeRepo.On("GetEmployeeNumberFromRFID", rfid, u.CompanyID).Return(employeeNumber, nil)

	employeeRepo.On("GetEmployeeAndBalance", employeeNumber, u.CompanyID).Return(emp, bal, nil)

	balAfter := bal.Balance - 200.00
	transactCheckoutMatcher := func(c entityCheckout.Checkout) bool {
		hasCorrectNetAmount := c.TotalNetAmount == 200
		hasCorrectNumberOfItems := len(c.CheckoutItems) == 1
		return hasCorrectNetAmount && hasCorrectNumberOfItems
	}
	transactBalMatcher := func(b entityEmployee.Balance) bool {
		hasCorrectBalance := b.Balance == 5000
		return hasCorrectBalance
	}
	checkoutRepo.
		On("Transact", mock.MatchedBy(transactCheckoutMatcher), mock.MatchedBy(transactBalMatcher), balAfter).
		Return(nil)

	// Execute create method
	actual, err := service.Create(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, 200, actual.StatusCode)

	type responseBody struct {
		Data restCheckout.CheckoutResponsePayload `json:"data"`
	}

	var actualBody responseBody
	err = json.Unmarshal([]byte(actual.Body), &actualBody)
	require.NoError(t, err)
	assert.Equal(t, 1, len(actualBody.Data.Checkout.CheckoutItems))
	assert.Equal(t, 200.00, actualBody.Data.Checkout.TotalGrossAmount)
	assert.Equal(t, 5000.00, actualBody.Data.Employee.BalanceBefore)
	assert.Equal(t, 4800.00, actualBody.Data.Employee.BalanceAfter)

	employeeRepo.AssertCalled(t, "GetEmployeeNumberFromRFID", rfid, u.CompanyID)
	employeeRepo.AssertCalled(t, "GetEmployeeAndBalance", employeeNumber, u.CompanyID)

	checkoutRepo.AssertCalled(t, "Transact", mock.MatchedBy(transactCheckoutMatcher), mock.MatchedBy(transactBalMatcher), balAfter)
}
