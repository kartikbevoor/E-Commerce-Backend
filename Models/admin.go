package models

type Admin struct {
	Id       int
	Name     string
	Age      int
	Email    string
	Username string
	Password string
	Address  Address
}
