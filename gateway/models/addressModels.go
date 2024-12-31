package models

type Address struct {
	AddressID     string `json:"addressId"`     // Unique identifier for the address
	UserID        string `json:"userId"`        // User ID associated with the address
	IsDefault     bool   `json:"isDefault"`     // Indicates if this is the default address
	Street        string `json:"street"`        // Street name or address line
	ApartmentUnit string `json:"apartmentUnit"` // Apartment or unit number
	City          string `json:"city"`          // City of the address
	State         string `json:"state"`         // State of the address
	Country       string `json:"country"`       // Country of the address
	ZipCode       string `json:"zipCode"`       // ZIP or postal code
	CreatedAt     string `json:"createdAt"`     // Timestamp of when the address was created
	Name          string `json:"name"`          // Name of the person associated with the address
	Phone         string `json:"phone"`         // Phone number associated with the address
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
