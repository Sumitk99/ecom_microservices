package models

type CartOpsReq struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  uint64 `json:"quantity" validate:"required,min=1"`
	CartName  string `json:"cart_name"`
}

type CartItem struct {
	ProductID string  `json:"product_id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
	SellerID  string  `json:"seller_id"`
	ImageURL  string  `json:"image_url"`
}

type CartResponse struct {
	CartName   string      `json:"cart_name"`
	Items      []*CartItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
}

type CheckoutRequest struct {
	CartID          string `json:"cart_id" validate:"required"`
	MethodOfPayment string `json:"method_of_payment" validate:"required"`
	TransactionID   string `json:"transaction_id"`
	AddressId       string `json:"address_id" validate:"required"`
}

type GetCartRequest struct {
	CartID string `json:"cart_id" validate:"required"`
}
