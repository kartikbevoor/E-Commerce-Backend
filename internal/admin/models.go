package admin

type Permissions string

const (
	RoleCustomer Permissions = "Level1"
	RoleVendor   Permissions = "Level2"
	RoleAdmin    Permissions = "Level3"
)

type Admin struct {
	UserID      int64
	Permissions Permissions
}
