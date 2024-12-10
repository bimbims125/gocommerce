package entity

type Cart struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	UserID    int `json:"user_id"`
}
