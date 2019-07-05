package db

type (
	// Product - makeup unit.
	Product struct {
		ID          int    `json:"product_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Apply       string `json:"apply"`
		Price       int    `json:"price"`
		Avatar      string `json:"avatar"`
		Brand       string `json:"brand_name"`
	}
	Brand struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
