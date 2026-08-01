package inventory

type CreateInventoryRequest struct {
	ProductID int `json:"product_id" validate:"gt=0"`
	Quantity  int `json:"quantity" validate:"gt=0"`
}

type CreateInventoryResponse struct {
	ID           int `json:"id"`
	ProductID    int `json:"product_id"`
	Quantity     int `json:"quantity"`
	SoldQuantity int `json:"sold_quantity"`
}
