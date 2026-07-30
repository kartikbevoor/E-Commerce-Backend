package vendors

type CreateVendorRequest struct {
	Name     string
	Age      int
	Email    string
	Username string
	Password string
}

type CreateVendorResponse struct {
	Name     string
	Age      int
	Email    string
	Username string
}
