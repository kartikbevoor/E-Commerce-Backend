package models

type Order struct {
	Id          int
	IsDelivered bool
	IsPaid      bool
	PaymentId   int
	CustomerId  int
	Note        string
	ProductsIds []int
	Amount      int
}
