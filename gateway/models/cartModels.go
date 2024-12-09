package models

type CartOpsReq struct {
	ProductID string `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
	CartName  string `json:"cart_name"`
}

type CartItem struct {
	ProductID string  `json:"product_id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
	SellerID  string  `json:"seller_id"`
}

type CartResponse struct {
	CartName   string      `json:"cart_name"`
	Items      []*CartItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
}
