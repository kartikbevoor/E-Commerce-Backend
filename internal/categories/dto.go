package categories

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Description string `json:"description" validate:"required,min=10,max=200"`
}

type CreateCategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
