package customer

type CreateCustomerRequest struct {
	UserId int64  `json:"user_id" validate:"required"`
	Gender string `json:"gender" validate:"required,oneof=male female other"`
	DOB    string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type CreateCustomerResponse struct {
	UserId        int64  `json:"user_id"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	LoyaltyPoints int64  `json:"loyalty_points"`
}
