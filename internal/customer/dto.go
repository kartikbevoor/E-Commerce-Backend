package customer

type CreateCustomerRequest struct {
	Name     string
	Age      int
	Email    string
	Username string
	Password string
	// Address  Address
}

type CreateCustomerResponse struct {
	Name     string
	Age      int
	Email    string
	Username string
	// Address  Address
}
