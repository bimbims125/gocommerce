package entity

type Order struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	ProductID     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	TotalPrice    float64 `json:"total_price"`
	TransactionID string  `json:"transaction_id"`
}

type GetOrder struct {
	ID            int          `json:"id"`
	User          UserOrder    `json:"user"`
	Product       ProductOrder `json:"product"`
	Quantity      int          `json:"quantity"`
	TotalPrice    float64      `json:"total_price"`
	TransactionID string       `json:"transaction_id"`
}

type ProductOrder struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type UserOrder struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
