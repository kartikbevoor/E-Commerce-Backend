package models

type Order struct {
	Id          int
	IsDelivered bool
	IsPaid      bool
	PaymentId   int
	CustomerId  int
	ProductsIds []int
}
