package vendors

type Vendor struct {
	UserID              int64  `db:"user_id"`
	StoreName           string `db:"store_name"`
	GstNumber           string `db:"gst_number"`
	PanNumber           string `db:"pan_number"`
	BusinessDescription string `db:"business_description"`
}
