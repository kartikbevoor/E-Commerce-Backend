package cart

import "time"

type Cart struct {
	Id          int
	CustomerId  int
	ProductsIds []int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
