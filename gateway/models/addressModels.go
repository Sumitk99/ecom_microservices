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

type AddAddressRequest struct {
	Phone         string `json:"phone" validate:"required"`
	IsDefault     bool   `json:"isDefault"`
	Street        string `json:"street" validate:"required"`
	ApartmentUnit string `json:"apartmentUnit"`
	City          string `json:"city" validate:"required"`
	State         string `json:"state" validate:"required"`
	Country       string `json:"country" validate:"required"`
	ZipCode       string `json:"zipCode" validate:"required"`
	Name          string `json:"name" validate:"required,min=2,max=100"`
}
