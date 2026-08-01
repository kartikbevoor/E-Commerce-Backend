package payments

type CreatePaymentRequest struct {
	Type       string `json:"type" validate:"required,oneof=card upi netbanking cod wallet"`
	CustomerID int    `json:"customer_id" validate:"gt=0"`
	Amount     int64  `json:"amount" validate:"gt=0"`
	OrderID    int    `json:"order_id" validate:"gt=0"`
}

type CreatePaymentResponse struct {
	Type       string `json:"type"`
	CustomerID int    `json:"customer_id"`
	Amount     int64  `json:"amount"`
	OrderID    int    `json:"order_id"`
}
