package user

import "time"

type Role string

const (
	RoleCustomer Role = "Customer"
	RoleVendor   Role = "Vendor"
	RoleAdmin    Role = "Admin"
)

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	Role         Role
	IsVerified   bool
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
