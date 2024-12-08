package helper

import (
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/cart/models"
	"github.com/Sumitk99/ecom_microservices/cart/pb"
	"github.com/Sumitk99/ecom_microservices/catalog"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/segmentio/ksuid"
	"log"
	"time"
)

const SECRET_KEY = "SECRET_KEY"

type SignedDetails struct {
	GuestId string
	jwt.StandardClaims
}

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

func GenerateGuestToken() (string, error) {
	claims := &SignedDetails{
		GuestId: ksuid.New().String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	singedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Println("Error during signing string", err)
		return "", errors.New("Error during signing string")
	}
	return singedToken, nil
}

func ValidateGuestToken(signedToken string) (string, error) {
	log.Println(signedToken)
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		log.Println(err)
		return "", errors.New(fmt.Sprintf("Invalid token: %v", err))
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		log.Println("Token is invalid")
		return "", errors.New("Token is invalid")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("Token is expired")
		return "", errors.New("Token is expired")
	}
	return claims.GuestId, nil
}
