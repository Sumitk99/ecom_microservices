package helper

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/cart/models"
	"github.com/Sumitk99/ecom_microservices/cart/pb"
	"github.com/Sumitk99/ecom_microservices/catalog"
)

func MakeProductArray(CartProducts []models.CartItem, IdToQuantity *map[string]uint64) *[]string {
	productIds := []string{}
	for _, item := range CartProducts {
		fmt.Printf("productId : %s, quantity : %d\n", item.ProductID, item.Quantity)
		productIds = append(productIds, item.ProductID)
		(*IdToQuantity)[item.ProductID] = item.Quantity
		fmt.Println("IdToQuantity : ", (*IdToQuantity)[item.ProductID])
	}
	return &productIds
}

func ProcessCart(products []catalog.Product, IdToQuantity map[string]uint64) ([]*pb.CartItem, float64) {
	var CartItems []*pb.CartItem
	totalPrice := 0.0
	for _, p := range products {
		CartItems = append(CartItems, &pb.CartItem{
			ProductId: p.ID,
			Title:     p.Name,
			Price:     p.Price,
			Quantity:  IdToQuantity[p.ID],
		})
		totalPrice += p.Price * float64(IdToQuantity[p.ID])
	}
	return CartItems, totalPrice
}
