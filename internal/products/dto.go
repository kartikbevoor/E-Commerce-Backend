package products

// Look into it figure out price and stuff

type CreateProductRequest struct {
	Name        string
	CategoryId  int
	VendorId    int
	Description string
}

type CreateProductResponsse struct {
	Name        string
	CategoryId  int
	VendorId    int
	Description string
}
