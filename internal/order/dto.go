package order

// Make clear objectives on products and their quantity

type CreateOrderRequest struct {
	CustomerId  int
	Note        string
	ProductsIds []int
	Amount      int
}

type CreateOrderResponse struct {
	CustomerId  int
	Note        string
	ProductsIds []int
	Amount      int
}
