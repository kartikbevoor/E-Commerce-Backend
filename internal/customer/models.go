package customer

import "time"

type Customer struct {
	Id        int
	Name      string
	Age       int
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Address  Address
}
