package payments

type CreatePaymentRequest struct {
	Type       string
	CustomerId int
	Amount     int
	OrderId    int
}

type CreatePaymentResponse struct {
	Type       string
	CustomerId int
	Amount     int
	OrderId    int
}
