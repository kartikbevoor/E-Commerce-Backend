package admin

import "time"

type Admin struct {
	Id        int
	Name      string
	Age       int
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	// Address  Address
}
