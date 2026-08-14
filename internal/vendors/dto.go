package vendors

type CreateVendorRequest struct {
	UserID              int64  `json:"user_id" validate:"required"`
	StoreName           string `json:"store_name" validate:"required"`
	GstNumber           string `json:"gst_number" validate:"required,len=15"`
	PanNumber           string `json:"pan_number" validate:"required,len=10"`
	BusinessDescription string `json:"business_description" validate:"required"`
}

type CreateVendorResponse struct {
	UserID              int64  `json:"user_id"`
	StoreName           string `json:"store_name"`
	GstNumber           string `json:"gst_number"`
	PanNumber           string `json:"pan_number"`
	BusinessDescription string `json:"business_description"`
}
