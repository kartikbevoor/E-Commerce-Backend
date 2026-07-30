package cart

type CreateCartRequest struct {
	CustomerId  int
	ProductsIds []int
}

type CreateCartResponse struct {
	CustomerId  int
	ProductsIds []int
}
