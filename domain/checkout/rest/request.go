package rest

type CreateRequestPayload struct {
	RFID    string                       `json:"rfid"`
	Station string                       `json:"station"`
	Items   []CheckoutItemRequestPayload `json:"items"`
}

type CheckoutItemRequestPayload struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Calories float64 `json:"calories"`
	Quantity int     `json:"qty"`
}
