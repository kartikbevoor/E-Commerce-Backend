package customer

type CreateCustomerRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=30"`
	Age      int    `json:"age" validate:"required,gte=18"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=72"`
	// Address  Address
}

type CreateCustomerResponse struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
	// Address  Address
}
