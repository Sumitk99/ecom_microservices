package models

import (
	"time"
)

type Card struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	CardNumber     string    `json:"card_number"`
	CardHolderName string    `json:"card_holder_name"`
	ExpiryMonth    string    `json:"expiry_month"`
	ExpiryYear     string    `json:"expiry_year"`
	CardType       string    `json:"card_type"`
	IsDefault      bool      `json:"is_default"`
	CreatedAt      time.Time `json:"created_at"`
}

type PaymentResult struct {
	TransactionID string
	Status        string
	Timestamp     time.Time
}
