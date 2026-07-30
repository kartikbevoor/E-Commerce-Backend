package reviews

type CreateReviewRequest struct {
	Msg        string
	CustomerId int
	ProductId  int
}

type CreateReviewResponse struct {
	Msg        string
	CustomerId int
	ProductId  int
}
