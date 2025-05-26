package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/Sumitk99/ecom_microservices/payment/models"
	"github.com/google/uuid"
)

type Service interface {
	AddCard(ctx context.Context, card *models.Card) (*models.Card, error)
	RemoveCard(ctx context.Context, cardID string) error
	ListCards(ctx context.Context, userID string) ([]*models.Card, error)
	ProcessPayment(ctx context.Context, userID, cardID string, amount float64, currency, description string) (*models.PaymentResult, error)
}

type PaymentService struct {
	repo Repository
}

func NewPaymentService(repo Repository) Service {
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) AddCard(ctx context.Context, card *models.Card) (*models.Card, error) {
	card.ID = uuid.New().String()
	card.CreatedAt = time.Now()

	err := s.repo.AddCard(card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *PaymentService) RemoveCard(ctx context.Context, cardID string) error {
	return s.repo.RemoveCard(cardID)
}

func (s *PaymentService) ListCards(ctx context.Context, userID string) ([]*models.Card, error) {
	return s.repo.ListCards(userID)
}

func (s *PaymentService) ProcessPayment(ctx context.Context, userID, cardID string, amount float64, currency, description string) (*models.PaymentResult, error) {
	// Fetch the card to ensure it belongs to the user and exists
	card, err := s.repo.GetCard(cardID)
	if err != nil {
		return nil, err
	}
	if card.UserID != userID {
		return nil, fmt.Errorf("unauthorized: card does not belong to user")
	}

	transactionID := uuid.New().String()

	result := &models.PaymentResult{
		TransactionID: transactionID,
		Status:        "success",
		Timestamp:     time.Now(),
	}

	err = s.repo.SaveTransaction(userID, cardID, amount, currency, description, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
