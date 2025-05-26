package server

import (
	"context"

	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
)

func (srv *Server) AddCard(card *models.AddCardRequest) (*pb.AddCardResponse, error) {
	res, err := srv.PaymentClient.AddCard(context.Background(), &pb.AddCardRequest{
		UserId:         card.UserID,
		CardNumber:     card.CardNumber,
		CardHolderName: card.CardHolderName,
		ExpiryMonth:    card.ExpiryMonth,
		ExpiryYear:     card.ExpiryYear,
		CardType:       card.CardType,
		IsDefault:      card.IsDefault,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (srv *Server) RemoveCard(req *models.RemoveCardRequest) (*pb.RemoveCardResponse, error) {
	res, err := srv.PaymentClient.RemoveCard(context.Background(), &pb.RemoveCardRequest{
		CardId: req.CardID,
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (srv *Server) ListCards(userID string) (*pb.ListCardsResponse, error) {
	res, err := srv.PaymentClient.ListCards(context.Background(), &pb.ListCardsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (srv *Server) ProcessPayment(req *models.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	res, err := srv.PaymentClient.ProcessPayment(context.Background(), &pb.ProcessPaymentRequest{
		UserId:      req.UserID,
		CardId:      req.CardID,
		Amount:      req.Amount,
		Currency:    req.Currency,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
