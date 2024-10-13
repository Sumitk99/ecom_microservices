package catalog

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("Product not found")
)

type Reposity interface {
	Close() error
	PutProduct(ctx context.Context, p Product) error
	GetProductByID(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	ListProductWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type elasticRepository struct {
	client *elastic.Client
}

type productDocument struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

func NewElasticRepository(url string) (Repository, error) {

}
