package inventory

// Before proceding make your objectives clear

type CreateInventoryRequest struct {
	Id          int
	IsDelivered bool
	IsPaid      bool
	PaymentId   int
	CustomerId  int
	Note        string
	ProductsIds []int
	Amount      int
}
