package order

import (
	"context"
	"github.com/segmentio/ksuid"
	"time"
)

type Service interface {
	PostOrder(ctx context.Context, accountID, MethodOfPayment, TransactionID, PaymentStatus string, products []OrderedProduct) (*Order, error)
	GetOrder(ctx context.Context, orderID, accountID string) (*Order, error)
	GetOrdersForAccount(ctx context.Context, accountID string) ([]*UserOrder, error)
}

type Order struct {
	ID              string
	CreatedAt       string
	TotalPrice      float64
	AccountID       string
	MethodOfPayment string
	TransactionID   string
	PaymentStatus   string
	ETA             string
	Products        []OrderedProduct
	OrderStatus     string
}

type OrderedProduct struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Quantity    uint32
	ImageURL    string
}

type orderService struct {
	repository Repository
}

type UserOrder struct {
	OrderId     string
	CreatedAt   string
	TotalPrice  string
	ETA         string
	OrderStatus string
}

func NewService(r Repository) Service {
	return &orderService{r}
}

func (s *orderService) PostOrder(ctx context.Context, accountID, MethodOfPayment, TransactionID, PaymentStatus string, products []OrderedProduct) (*Order, error) {
	order := &Order{
		ID:              ksuid.New().String(),
		CreatedAt:       time.Now().String(),
		AccountID:       accountID,
		Products:        products,
		MethodOfPayment: MethodOfPayment,
		TransactionID:   TransactionID,
		PaymentStatus:   PaymentStatus,
	}

	order.TotalPrice = 0.0
	for _, p := range products {
		order.TotalPrice += p.Price * float64(p.Quantity)
	}

	err := s.repository.PutOrder(ctx, *order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderService) GetOrder(ctx context.Context, orderID, accountID string) (*Order, error) {
	order, err := s.repository.GetOrder(ctx, orderID, accountID)
	if err != nil {
		return nil, err
	}
	return order, err
}

func (s *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]*UserOrder, error) {
	return s.repository.GetOrdersForAccount(ctx, accountID)
}
