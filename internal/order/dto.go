package order

type CreateOrderRequest struct {
	ProductID  int    `json:"product_id" validate:"gt=0"`
	CustomerID int    `json:"customer_id" validate:"gt=0"`
	Quantity   int    `json:"quantity" validate:"gt=0"`
	Note       string `json:"note"`
}

type CreateOrderResponse struct {
	ProductID   int   `json:"product_id"`
	CustomerID  int   `json:"customer_id"`
	Quantity    int   `json:"quantity"`
	Amount      int64 `json:"amount"` // stored in paisa
	IsDelivered bool  `json:"is_delivered"`
	IsPaid      bool  `json:"is_paid"`
	PaymentID   int   `json:"payment_id"`
}
