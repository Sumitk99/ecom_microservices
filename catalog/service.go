package catalog

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/catalog/models"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostProduct(ctx context.Context, name, description, sellerID, sellerName, imageUrl, category string, price float64, stock uint64, locations, sizes []string, colors []models.Color) (*models.Product, error)
	GetProduct(ctx context.Context, id string) (*models.Product, error)
	GetProducts(ctx context.Context, skip, take uint64) ([]models.Product, error)
	GetProductByIDs(ctx context.Context, ids []string) ([]models.Product, error)
	SearchProducts(ctx context.Context, query string, skip, take uint64) ([]models.Product, error)
}

type catalogService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &catalogService{r}
}

func (s *catalogService) PostProduct(ctx context.Context, name, description, sellerID, sellerName, imageUrl, category string, price float64, stock uint64, locations, sizes []string, colors []models.Color) (*models.Product, error) {
	p := models.Product{
		ID:          ksuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		SellerID:    sellerID,
		SellerName:  sellerName,
		ImageUrl:    imageUrl,
		Category:    category,
		Stock:       stock,
		Locations:   locations,
		Sizes:       sizes,
		Colors:      colors,
	}
	if err := s.repository.PutProduct(ctx, p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *catalogService) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	return s.repository.GetProductByID(ctx, id)
}

func (s *catalogService) GetProducts(ctx context.Context, skip, take uint64) ([]models.Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.ListProducts(ctx, skip, take)
}

func (s *catalogService) GetProductByIDs(ctx context.Context, ids []string) ([]models.Product, error) {
	return s.repository.ListProductWithIDs(ctx, ids)
}

func (s *catalogService) SearchProducts(ctx context.Context, query string, skip, take uint64) ([]models.Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.SearchProducts(ctx, query, skip, take)
}
