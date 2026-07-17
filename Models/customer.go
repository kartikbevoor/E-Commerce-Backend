package models

type Customer struct {
	Id       int
	Name     string
	Age      int
	Email    string
	Username string
	Password string
	Address  Address
}
