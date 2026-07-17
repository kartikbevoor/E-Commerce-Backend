package models

type Vendor struct {
	Id       int
	Name     string
	Age      int
	Email    string
	Username string
	Password string
	Address  Address
	Products []Products
}
