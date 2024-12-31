package models

type Address struct {
	AddressID     string `json:"addressId"`
	UserID        string `json:"userId"`
	IsDefault     bool   `json:"isDefault"`
	Street        string `json:"street"`
	ApartmentUnit string `json:"apartmentUnit"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zipCode"`
	CreatedAt     string `json:"createdAt"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
}

type Order struct {
	ID              string
	CreatedAt       string
	TotalPrice      float64
	AccountID       string
	MethodOfPayment string
	TransactionID   string
	PaymentStatus   string
	ETA             string
	Products        []OrderedProduct
	OrderStatus     string
	AddressId       string
}

type OrderedProduct struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Quantity    uint32
	ImageURL    string
}

type UserOrder struct {
	OrderId     string
	CreatedAt   string
	TotalPrice  string
	ETA         string
	OrderStatus string
}
