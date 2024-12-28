package models

type Color struct {
	ColorName string `json:"color_name"`
	Hex       string `json:"hex"`
}

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	SellerID    string   `json:"seller_id"`
	SellerName  string   `json:"seller_name"`
	ImageUrl    string   `json:"image_url"`
	Category    string   `json:"category"`
	Stock       uint64   `json:"stock"`
	Locations   []string `json:"locations"`
	Sizes       []string `json:"sizes"`
	Colors      []Color  `json:"colors"`
}

type ProductDocument struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	SellerID    string   `json:"seller_id"`
	SellerName  string   `json:"seller_name"`
	ImageUrl    string   `json:"image_url"`
	Category    string   `json:"category"`
	Stock       uint64   `json:"stock"`
	Locations   []string `json:"locations"`
	Sizes       []string `json:"sizes"`
	Colors      []Color  `json:"colors"`
}
