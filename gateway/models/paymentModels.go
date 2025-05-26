package models

type AddCardRequest struct {
	UserID         string `json:"user_id"`
	CardNumber     string `json:"card_number"`
	CardHolderName string `json:"card_holder_name"`
	ExpiryMonth    string `json:"expiry_month"`
	ExpiryYear     string `json:"expiry_year"`
	CardType       string `json:"card_type"`
	IsDefault      bool   `json:"is_default"`
}

type RemoveCardRequest struct {
	UserID string `json:"user_id"`
	CardID string `json:"card_id"`
}

type ProcessPaymentRequest struct {
	UserID      string  `json:"user_id"`
	CardID      string  `json:"card_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
}
