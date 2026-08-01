package cart

type CreateCartRequest struct {
	CustomerID  int   `json:"customer_id" validate:"required"`
	ProductsIDs []int `json:"product_ids" validate:"required,min=1,dive,gt=0"` // here: the dive applies the validation to each following validation
}

type CreateCartResponse struct {
	CustomerID  int   `json:"customer_id"`
	ProductsIDs []int `json:"product_ids"`
}
