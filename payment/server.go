package payment

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/payment/models"
	"github.com/Sumitk99/ecom_microservices/payment/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type PaymentServer struct {
	service Service
	pb.UnimplementedPaymentServiceServer
}

func ListenGRPC(s Service, port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterPaymentServiceServer(srv, &PaymentServer{
		service:                           s,
		UnimplementedPaymentServiceServer: pb.UnimplementedPaymentServiceServer{},
	})
	reflection.Register(srv)
	return srv.Serve(lis)
}

func (s *PaymentServer) AddCard(ctx context.Context, req *pb.AddCardRequest) (*pb.AddCardResponse, error) {
	card := &models.Card{
		UserID:         req.UserId,
		CardNumber:     req.CardNumber,
		CardHolderName: req.CardHolderName,
		ExpiryMonth:    req.ExpiryMonth,
		ExpiryYear:     req.ExpiryYear,
		CardType:       req.CardType,
		IsDefault:      req.IsDefault,
	}

	res, err := s.service.AddCard(ctx, card)
	if err != nil || res == nil {
		return &pb.AddCardResponse{Error: err.Error()}, err
	}

	return &pb.AddCardResponse{
		Card: &pb.Card{
			Id:             res.ID,
			UserId:         res.UserID,
			CardNumber:     res.CardNumber,
			CardHolderName: res.CardHolderName,
			ExpiryMonth:    res.ExpiryMonth,
			ExpiryYear:     res.ExpiryYear,
			CardType:       res.CardType,
			IsDefault:      res.IsDefault,
			CreatedAt:      res.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *PaymentServer) RemoveCard(ctx context.Context, req *pb.RemoveCardRequest) (*pb.RemoveCardResponse, error) {
	err := s.service.RemoveCard(ctx, req.CardId)
	if err != nil {
		return &pb.RemoveCardResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}
	return &pb.RemoveCardResponse{
		Success: true,
	}, nil
}

func (s *PaymentServer) ListCards(ctx context.Context, req *pb.ListCardsRequest) (*pb.ListCardsResponse, error) {

	cards, err := s.service.ListCards(ctx, req.UserId)
	if err != nil {
		return &pb.ListCardsResponse{
			Error: err.Error(),
		}, err
	}

	var pbCards []*pb.Card
	for _, c := range cards {
		pbCards = append(pbCards, &pb.Card{
			Id:             c.ID,
			UserId:         c.UserID,
			CardNumber:     c.CardNumber,
			CardHolderName: c.CardHolderName,
			ExpiryMonth:    c.ExpiryMonth,
			ExpiryYear:     c.ExpiryYear,
			CardType:       c.CardType,
			IsDefault:      c.IsDefault,
			CreatedAt:      c.CreatedAt.Format(time.RFC3339),
		})
	}

	return &pb.ListCardsResponse{
		Cards: pbCards,
	}, nil
}

func (s *PaymentServer) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	result, err := s.service.ProcessPayment(ctx, req.UserId, req.CardId, req.Amount, req.Currency, req.Description)
	if err != nil {
		return &pb.ProcessPaymentResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.ProcessPaymentResponse{
		TransactionId: result.TransactionID,
		Status:        result.Status,
	}, nil
}
