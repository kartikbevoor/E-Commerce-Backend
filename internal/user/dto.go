package user

import "time"

type CreateUserRequest struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validare:"email"`
	PasswordHash string `json:"password_hash" vali`
	Role         Role
}

type CreateUserResponse struct {
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
