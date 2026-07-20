package models

type Review struct {
	Id         int
	Msg        string
	CustomerId int
	ProductId  int
}
