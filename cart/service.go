package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/cart/helper"
	"github.com/Sumitk99/ecom_microservices/cart/models"
	"google.golang.org/grpc/metadata"
	"log"
)

type CartService interface {
	AddItem(ctx context.Context, productId string, quantity uint64) error
	DeleteItem(ctx context.Context, productId string) error
	GetCartItems(ctx context.Context) ([]models.CartItem, error)
	UpdateItem(ctx context.Context, productId string, quantity uint64) error
	DeleteCart(ctx context.Context) error
	IssueGuestToken(ctx context.Context) (string, error)
	ValidateGuestId(ctx context.Context) (string, error)
}

type cartService struct {
	repository Repository
}

func NewService(r Repository) CartService {
	return &cartService{r}
}

func (s *cartService) AddItem(ctx context.Context, productId string, quantity uint64) error {
	if len(productId) == 0 || quantity <= 0 {
		return errors.New("Invalid Input")
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return errors.New("no user metadata found in context")
	}
	var emptyString string
	account, cart, guestId := md.Get("UserID"), md.Get("CartID"), md.Get("GuestID")
	fmt.Println(account)
	fmt.Println(cart)
	fmt.Println(guestId)
	var err error
	fmt.Printf("%s %s %s\n", account[0], cart[0], guestId[0])
	if len(account[0]) > 0 && len(cart[0]) > 0 {
		err = s.repository.AddItem(ctx, cart[0], account[0], emptyString, productId, quantity)
	} else if len(guestId[0]) > 0 {
		fmt.Println("guestSection Included")
		err = s.repository.AddItem(ctx, emptyString, emptyString, guestId[0], productId, quantity)
	} else {
		return errors.New("not Enough Data to Insert Item to Cart")
	}
	return err
}

func (s *cartService) GetCartItems(ctx context.Context) ([]models.CartItem, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}
	var emptyString string
	account, cart, guestId := md.Get("UserID"), md.Get("CartID"), md.Get("GuestID")

	if len(account[0]) > 0 && len(cart[0]) > 0 {
		return s.repository.GetCartItems(ctx, cart[0], account[0], emptyString)
	} else if len(guestId[0]) > 0 {
		return s.repository.GetCartItems(ctx, emptyString, emptyString, guestId[0])
	} else {
		return nil, errors.New("not Enough Data to Insert Item to Cart")
	}

}

func (s *cartService) DeleteItem(ctx context.Context, productId string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return errors.New("no user metadata found in context")
	}
	var emptyString string
	account, cart, guestId := md.Get("UserID"), md.Get("CartID"), md.Get("GuestID")
	var err error
	if len(account[0]) > 0 && len(cart[0]) > 0 {
		err = s.repository.DeleteItem(ctx, cart[0], account[0], emptyString, productId)
	} else if len(guestId[0]) > 0 {
		err = s.repository.DeleteItem(ctx, emptyString, emptyString, guestId[0], productId)
	} else {
		return errors.New("not Enough Data to Insert Item to Cart")
	}
	return err
}

func (s *cartService) UpdateItem(ctx context.Context, productId string, quantity uint64) error {
	if quantity <= 0 {
		return errors.New("Min Quantity is 1, Select Delete Item to remove item from Cart")
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return errors.New("no user metadata found in context")
	}
	var emptyString string
	account, cart, guestId := md.Get("UserID"), md.Get("CartID"), md.Get("GuestID")
	var err error

	if len(account[0]) > 0 && len(cart[0]) > 0 {
		err = s.repository.UpdateItem(ctx, cart[0], account[0], emptyString, productId, quantity)
	} else if len(guestId[0]) > 0 {
		err = s.repository.UpdateItem(ctx, emptyString, emptyString, guestId[0], productId, quantity)
	} else {
		return errors.New("not Enough Data to Insert Item to Cart")
	}
	return err
}

func (s *cartService) DeleteCart(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return errors.New("no user metadata found in context")
	}
	var emptyString string
	account, cart, guestId := md.Get("UserID"), md.Get("CartID"), md.Get("GuestID")
	var err error
	if len(account[0]) > 0 && len(cart[0]) > 0 {
		err = s.repository.DeleteCart(ctx, cart[0], account[0], emptyString)
	} else if len(guestId[0]) > 0 {
		err = s.repository.DeleteCart(ctx, emptyString, emptyString, guestId[0])
	} else {
		return errors.New("not Enough Data to Insert Item to Cart")
	}
	return err
}

func (s *cartService) IssueGuestToken(ctx context.Context) (string, error) {
	return helper.GenerateGuestToken()
}

func (s *cartService) ValidateGuestId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return "", errors.New("no user metadata found in context")
	}
	guestToken := md.Get("guestToken")
	if len(guestToken) == 0 {
		return "", errors.New("no guest token found in context")
	}

	guestId, err := helper.ValidateGuestToken(guestToken[0])
	if err != nil {
		return "", err
	}
	return guestId, err
}
