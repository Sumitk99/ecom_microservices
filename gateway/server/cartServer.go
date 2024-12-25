package server

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) AddItemToCart(ctx context.Context, productId string, quantity uint64) (*models.CartResponse, error) {
	res, err := s.CartClient.AddItemToCart(ctx, &pb.AddToCartRequest{
		ProductId: productId,
		Quantity:  quantity,
	})
	fmt.Println("Grpc response received")

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ProcessCartResponse(res), err
}

func (s *Server) ValidateGuestCartToken(ctx context.Context, token string) (string, error) {
	md := metadata.New(map[string]string{
		"guestToken": token,
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	res, err := s.CartClient.ValidateGuestCartToken(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return res.GuestId, nil
}

func (s *Server) IssueGuestCartToken(ctx context.Context) (string, error) {
	res, err := s.CartClient.IssueGuestCartToken(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return res.GuestToken, err
}

func (s *Server) GetCart(ctx context.Context) (*models.CartResponse, error) {

	res, err := s.CartClient.GetCart(ctx, &emptypb.Empty{})
	fmt.Println("Grpc response received")

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ProcessCartResponse(res), err
}

func (s *Server) RemoveItemFromCart(ctx context.Context, productId string) (*models.CartResponse, error) {
	res, err := s.CartClient.RemoveItemFromCart(ctx, &pb.RemoveFromCartRequest{
		ProductId: productId,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ProcessCartResponse(res), err
}

func (s *Server) UpdateCart(ctx context.Context, productId string, updatedQuantity uint64) (*models.CartResponse, error) {
	res, err := s.CartClient.UpdateCart(ctx, &pb.UpdateCartRequest{
		ProductId:       productId,
		UpdatedQuantity: updatedQuantity,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ProcessCartResponse(res), err
}

func (s *Server) DeleteCart(ctx context.Context) error {
	_, err := s.CartClient.DeleteCart(ctx, &emptypb.Empty{})
	return err
}

func ProcessCartResponse(res *pb.CartResponse) *models.CartResponse {
	cart := models.CartResponse{
		CartName:   res.Cart.CartId,
		TotalPrice: res.Cart.TotalPrice,
	}

	for _, item := range res.Cart.Items {
		cart.Items = append(cart.Items, &models.CartItem{
			ProductID: item.ProductId,
			Price:     item.Price,
			Quantity:  item.Quantity,
			Title:     item.Title,
		})
	}
	return &cart
}

func (s *Server) Checkout(ctx context.Context, cartId, methodOfPayment, transactionId string) (*models.Order, error) {
	log.Printf("Checking out cart %s with method of payment %s and transaction id %s", cartId, methodOfPayment, transactionId)
	res, err := s.CartClient.Checkout(ctx, &pb.CheckoutRequest{
		CartId:          cartId,
		MethodOfPayment: methodOfPayment,
		TransactionId:   transactionId,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	Order := &models.Order{
		OrderID:         res.Order.Id,
		MethodOfPayment: res.Order.MethodOfPayment,
		TransactionID:   res.Order.TransactionId,
		PaymentStatus:   res.Order.PaymentStatus,
		CreatedAt:       res.Order.CreatedAt,
		TotalPrice:      res.Order.TotalPrice,
		ETA:             res.Order.ETA,
		OrderStatus:     res.Order.OrderStatus,
		Products:        []*models.OrderedProduct{},
	}
	for _, product := range res.Order.Products {
		Order.Products = append(Order.Products, &models.OrderedProduct{
			ID:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	return Order, err
}
