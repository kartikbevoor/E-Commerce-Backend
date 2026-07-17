package models

type Cart struct {
	Id          int
	CustomerId  int
	ProductsIds []int
}
