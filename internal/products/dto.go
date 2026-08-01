package products

// Look into it figure out price and stuff

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	CategoryID  int    `json:"category_id" validate:"gt=0"`
	VendorID    int    `json:"vendor_id" validate:"gt=0"`
	Price       int64  `json:"price" validate:"gt=0"`
	Description string `json:"description" validate:"max=500"`
}

type CreateProductResponse struct {
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	VendorID   int    `json:"vendor_id"`
	Price      int64  `json:"price"`
}
