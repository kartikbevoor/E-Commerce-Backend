package address

import "time"

type Address struct {
	ID          int
	UserID      int
	Street      string
	City        string
	State       string
	Country     string
	Pincode     string
	AddressType string
	IsDefault   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
