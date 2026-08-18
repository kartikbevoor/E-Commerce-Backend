package customer

type Customer struct {
	UserID        int64  `db:"user_id"`
	LoyaltyPoints int64  `db:"loyalty_points"`
	DOB           string `db:"dob"`
	Gender        string `db:"gender"`
}
