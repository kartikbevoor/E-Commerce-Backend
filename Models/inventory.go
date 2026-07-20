package models

type Inventory struct {
	Id                int
	ProducID          int
	AvailableQuantity float64
	SoldQuantity      float64
}
