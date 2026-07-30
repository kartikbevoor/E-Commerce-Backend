package wishlist

type CreateWishlistRequest struct {
	CustomerId  int
	ProductsIds []int
}

type CreateWishlistResponse struct {
	CustomerId  int
	ProductsIds []int
}
