package cart

import (
	"context"
	"errors"
	"github.com/Sumitk99/ecom_microservices/cart/models"
	"google.golang.org/grpc/metadata"
	"log"
)

type CartService interface {
	AddItem(ctx context.Context, productId string, quantity int) error
	DeleteItem(ctx context.Context, productId string) error
	GetCartItems(ctx context.Context) ([]models.CartItem, error)
	UpdateItem(ctx context.Context, productId string, quantity uint64) error
	DeleteCart(ctx context.Context) error
}

type cartService struct {
	repository Repository
}

func NewService(r Repository) CartService {
	return &cartService{r}
}

func (s *cartService) AddItem(ctx context.Context, productId string, quantity int) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return errors.New("no user metadata found in context")
	}
	var emptyString string
	account, cart, guestId := md.Get("UserID"), md.Get("CartID"), md.Get("GuestID")
	var err error
	if account != nil && cart != nil {
		err = s.repository.AddItem(ctx, cart[0], account[0], emptyString, productId, quantity)
	} else if guestId != nil {
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

	if account != nil && cart != nil {
		return s.repository.GetCartItems(ctx, cart[0], account[0], emptyString)
	} else if guestId != nil {
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
	if account != nil && cart != nil {
		err = s.repository.DeleteItem(ctx, cart[0], account[0], emptyString, productId)
	} else if guestId != nil {
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

	if account != nil && cart != nil {
		err = s.repository.UpdateItem(ctx, cart[0], account[0], emptyString, productId, quantity)
	} else if guestId != nil {
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
	if account != nil && cart != nil {
		err = s.repository.DeleteCart(ctx, cart[0], account[0], emptyString)
	} else if guestId != nil {
		err = s.repository.DeleteCart(ctx, emptyString, emptyString, guestId[0])
	} else {
		return errors.New("not Enough Data to Insert Item to Cart")
	}
	return err
}
