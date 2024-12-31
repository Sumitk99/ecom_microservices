package server

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"log"
)

func (s *Server) GetOrder(ctx context.Context, orderId string) (*pb.GetOrderResponse, error) {
	res, err := s.OrderClient.GetOrder(ctx, &pb.GetOrderRequest{
		OrderId: orderId,
	})
	if err != nil {
		return nil, err
	}
	log.Println(res.Order.Address)

	return res, nil
}

func (s *Server) GetOrders(ctx context.Context, accountID string) (*pb.GetOrdersForAccountResponse, error) {
	res, err := s.OrderClient.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{
		AccountId: accountID,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
