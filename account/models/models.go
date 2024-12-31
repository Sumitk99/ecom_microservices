package models

import "time"

type Account struct {
	ID           string    `json:"id"`
	Name         string    `json:"name" validate:"required,min=2,max=100"`
	Password     string    `json:"password" validate:"required,min=6"`
	Email        string    `json:"email" validate:"email,required"`
	Phone        string    `json:"phone" validate:"required"`
	Token        string    `json:"token"`
	UserType     string    `json:"user_type" validate:"required,eq=ADMIN|eq=BUYER|eq=SELLER"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

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

// AddAddressRequest represents a request to add a new address.
type AddAddressRequest struct {
	Phone         string `json:"phone" validate:"required"`              // Phone number for the new address
	IsDefault     bool   `json:"isDefault"`                              // Indicates if the new address should be the default
	Street        string `json:"street" validate:"required"`             // Street name or address line
	ApartmentUnit string `json:"apartmentUnit"`                          // Apartment or unit number
	City          string `json:"city" validate:"required"`               // City of the new address
	State         string `json:"state" validate:"required"`              // State of the new address
	Country       string `json:"country" validate:"required"`            // Country of the new address
	ZipCode       string `json:"zipCode" validate:"required"`            // ZIP or postal code
	Name          string `json:"name" validate:"required,min=2,max=100"` // Name of the person associated with the new address
}
