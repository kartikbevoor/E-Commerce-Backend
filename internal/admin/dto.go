package admin

type CreateAdminRequest struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"gte=18"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,,min=3,max=30,alphanum"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateAdminResponse struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
