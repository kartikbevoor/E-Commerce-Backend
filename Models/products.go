package models

type Products struct {
	Id          int
	Name        string
	CategoryId  int
	VendorId    int
	Description string
	ReviewsIds  []int
}
