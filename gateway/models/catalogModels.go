package models

type Color struct {
	ColorName string `json:"color_name" form:"color_name"`
	Hex       string `json:"hex" form:"hex"`
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
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	Price       float64  `json:"price" form:"price"`
	SellerID    string   `json:"seller_id" form:"seller_id"`
	SellerName  string   `json:"seller_name" form:"seller_name"`
	ImageUrl    string   `json:"image_url" form:"image_url"`
	Category    string   `json:"category" form:"category"`
	Stock       uint64   `json:"stock" form:"stock"`
	Locations   []string `json:"locations" form:"locations"`
	Sizes       []string `json:"sizes" form:"sizes"`
	Colors      []Color  `json:"colors" form:"colors"`
}

type GetProductsRequest struct {
	Skip  uint64   `json:"skip"`
	Take  uint64   `json:"take"`
	Ids   []string `json:"ids"`
	Query string   `json:"query"`
}

type Products struct {
	Title      string  `json:"title"`
	ProductId  string  `json:"productId"`
	Price      float64 `json:"price"`
	SellerName string  `json:"sellerName"`
	ImageURL   string  `json:"imageURL"`
}

type GetProductsResponse struct {
	Products []Products `json:"products"`
}
