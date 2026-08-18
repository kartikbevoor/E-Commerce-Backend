package user

import "time"

type Role string

const (
	RoleCustomer Role = "customer"
	RoleVendor   Role = "vendor"
	RoleAdmin    Role = "admin"
)

type User struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Username   string    `db:"username"`
	Password   string    `db:"password"`
	Role       Role      `db:"role"`
	IsVerified bool      `db:"is_verified"`
	IsActive   bool      `db:"is_active"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
