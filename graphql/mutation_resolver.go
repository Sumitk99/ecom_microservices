package main

import (
	"context"
	"errors"
	"github.com/Sumitk99/ecom_microservices/order"
	"log"
	"time"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, in AccountInput) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	account, err := r.server.accountClient.PostAccount(ctx, in.Name)
	if err != nil {
		log.Println("Error Creating Account at Graphql ", err)
		return nil, err
	}

	return &Account{
		ID:   account.ID,
		Name: account.Name,
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, in ProductInput) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	product, err := r.server.catalogClient.PostProduct(ctx, in.Name, in.Description, in.Price)
	if err != nil {
		log.Println("Error Creating Product at Graphql ", err)
		return nil, err
	}

	return &Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, in OrderInput) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var products []order.OrderedProduct

	for _, p := range in.Products {
		if p.Quantity <= 0 {
			return nil, errors.New("Invalid parameter for Order")
		}

		products = append(products, order.OrderedProduct{
			ID:       p.ID,
			Quantity: uint32(p.Quantity),
		})
	}

	o, err := r.server.orderClient.PostOrder(ctx, in.AccountID, products)
	if err != nil {
		log.Println("Error Creating Product at Graphql ", err)
		return nil, err
	}
	return &Order{
		ID:         o.ID,
		CreatedAt:  o.CreatedAt,
		TotalPrice: o.TotalPrice,
		Products:   []*Product{},
	}, nil
}
