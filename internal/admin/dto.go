package admin

type CreateAdminRequest struct {
	UserID      int64       `json:"user_id" validate:"required"`
	Permissions Permissions `json:"permissions" validate:"required"`
}

type CreateAdminResponse struct {
	UserID      int64       `json:"user_id"`
	Permissions Permissions `json:"permissions"`
}
