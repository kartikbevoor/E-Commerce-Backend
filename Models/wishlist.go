package models

type Wishlist struct {
	Id          int
	CustomerId  int
	ProductsIds []int
}
