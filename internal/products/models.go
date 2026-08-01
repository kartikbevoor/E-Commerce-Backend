package products

type Products struct {
	ID          int
	Name        string
	CategoryID  int
	VendorID    int
	Price       int64
	Description string
	ReviewsIDs  []int
}
