package entity

type MidtransRequest struct {
	UserID   int    `json:"user_id"`
	Amount   int64  `json:"amount"`
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

type MidtransResponse struct {
	TransactionID string `json:"transaction_id"`
	Token         string `json:"token"`
	RedirectURL   string `json:"redirect_url"`
}
