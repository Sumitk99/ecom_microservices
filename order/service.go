package order

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/order/models"
	"github.com/segmentio/ksuid"
	"time"
)

type Service interface {
	PostOrder(ctx context.Context, accountID, MethodOfPayment, TransactionID, PaymentStatus, addressId string, products []models.OrderedProduct) (*models.Order, error)
	GetOrder(ctx context.Context, orderID, accountID string) (*models.Order, error)
	GetOrdersForAccount(ctx context.Context, accountID string) ([]*models.UserOrder, error)
}

type orderService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &orderService{r}
}

func (s *orderService) PostOrder(ctx context.Context, accountID, MethodOfPayment, TransactionID, PaymentStatus, addressId string, products []models.OrderedProduct) (*models.Order, error) {
	order := &models.Order{
		ID:              ksuid.New().String(),
		CreatedAt:       time.Now().String(),
		AccountID:       accountID,
		Products:        products,
		MethodOfPayment: MethodOfPayment,
		TransactionID:   TransactionID,
		PaymentStatus:   PaymentStatus,
		AddressId:       addressId,
	}

	order.TotalPrice = 0.0
	for _, p := range products {
		order.TotalPrice += p.Price * float64(p.Quantity)
	}

	err := s.repository.PutOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderService) GetOrder(ctx context.Context, orderID, accountID string) (*models.Order, error) {
	order, err := s.repository.GetOrder(ctx, orderID, accountID)
	if err != nil {
		return nil, err
	}
	return order, err
}

func (s *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]*models.UserOrder, error) {
	return s.repository.GetOrdersForAccount(ctx, accountID)
}
