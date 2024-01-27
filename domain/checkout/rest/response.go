package rest

import "github.com/pevin/internal-pos-service-api/domain/checkout/entity"

type EmployeeResponsePayload struct {
	EmployeeNumber string  `json:"employee_number"`
	Name           string  `json:"name"`
	ImageUrl       string  `json:"image_url"`
	BalanceBefore  float64 `json:"balance_before"`
	BalanceAfter   float64 `json:"balance_after"`
}

type CheckoutResponsePayload struct {
	Checkout entity.Checkout         `json:"checkout"`
	Employee EmployeeResponsePayload `json:"employee"`
}
