package payments

type Payments struct {
	ID         int
	Type       string
	CustomerID int
	Amount     int64
	OrderId    int
}
