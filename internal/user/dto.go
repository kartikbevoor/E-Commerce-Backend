package user

import "time"

type CreateUserRequest struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"email"`
	UserName     string `json:"username" validate:"required,min=4,max=30"`
	PasswordHash string `json:"password_hash" validate:"required"`
	Role         Role   `json:"role" validate:"required"`
}

type CreateUserResponse struct {
	ID         int       `json:"id"`
	Name       string    `josn:"name"`
	Email      string    `json:"email"`
	Role       Role      `json:"role"`
	IsVerified bool      `json:"is_verified"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
