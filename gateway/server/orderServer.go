package server

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"log"
)

func (s *Server) GetOrder(ctx context.Context, orderId string) (*models.Order, error) {
	res, err := s.OrderClient.GetOrder(ctx, &pb.GetOrderRequest{
		OrderId: orderId,
	})
	if err != nil {
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

func (s *Server) GetOrders(ctx context.Context, accountID string) (*models.UserOrders, error) {
	res, err := s.OrderClient.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{
		AccountId: accountID,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	orders := models.UserOrders{}
	for _, order := range res.Orders {
		orders.Orders = append(orders.Orders, &models.UserOrder{
			OrderId:     order.OrderId,
			CreatedAt:   order.CreatedAt,
			TotalPrice:  order.TotalPrice,
			ETA:         order.ETA,
			OrderStatus: order.OrderStatus,
		})
	}
	return &orders, nil
}
