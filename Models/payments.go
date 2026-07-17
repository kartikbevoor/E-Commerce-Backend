package models

type Payments struct {
	Id         int
	Type       string
	CustomerId int
	Amount     int
	OrderId    int
}
