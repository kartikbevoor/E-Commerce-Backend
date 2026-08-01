package order

type Order struct {
	Id          int
	ProductID   int
	CustomerId  int
	Quantity    int
	Amount      int64
	IsDelivered bool
	IsPaid      bool
	PaymentId   int
	Note        string
}
