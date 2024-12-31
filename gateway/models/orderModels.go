package models

type OrderedProduct struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint32  `json:"quantity"`
	ImageUrl string  `json:"image_url"`
}

type Order struct {
	Products        []*OrderedProduct `json:"products"`
	OrderID         string            `json:"order_id"`
	CreatedAt       string            `json:"created_at"`
	TotalPrice      float64           `json:"total_price"`
	ETA             string            `json:"eta"`
	MethodOfPayment string            `json:"method_of_payment"`
	TransactionID   string            `json:"transaction_id"`
	PaymentStatus   string            `json:"payment_status"`
	OrderStatus     string            `json:"order_status"`
	AddressID       Address           `json:"address_id"`
}

type UserOrder struct {
	OrderId     string `json:"order_id"`
	CreatedAt   string `json:"created_at"`
	TotalPrice  string `json:"total_price"`
	ETA         string `json:"eta"`
	OrderStatus string `json:"order_status"`
}

type UserOrders struct {
	Orders []*UserOrder `json:"orders"`
}
