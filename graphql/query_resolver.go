package main

import (
	"context"
	"log"
	"time"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if id != nil {
		r, err := r.server.accountClient.GetAccount(ctx, *id)
		if err != nil {
			log.Println("Error Getting Account at Graphql ", err)
			return nil, err
		}
		return []*Account{{
			ID:   r.ID,
			Name: r.Name,
		}}, nil
	}

	skip, take := uint64(0), uint64(0)
	skip, take = uint64(pagination.Skip), uint64(pagination.Take)

	accountList, err := r.server.accountClient.GetAccounts(ctx, skip, take)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var accounts []*Account
	for _, a := range accountList {
		accounts = append(accounts, &Account{
			ID:   a.ID,
			Name: a.Name,
		})
	}
	return accounts, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination PaginationInput, query, id *string) ([]*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if id != nil {
		res, err := r.server.catalogClient.GetProduct(ctx, *id)
		if err != nil {
			log.Println("Error Getting Account at Graphql ", err)
			return nil, err
		}
		return []*Product{{
			ID:          res.ID,
			Name:        res.Name,
			Description: res.Description,
			Price:       res.Price,
		}}, nil
	}

	skip, take := pagination.bounds()
	//skip, take = uint64(pagination.Skip), uint64(pagination.Take)

	q := ""
	if query != nil {
		q = *query
	}
	productList, err := r.server.catalogClient.GetProducts(ctx, skip, take, nil, q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var products []*Product
	for _, p := range productList {
		products = append(products, &Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return products, nil
}

func (p PaginationInput) bounds() (uint64, uint64) {
	skip, take := uint64(0), uint64(0)
	if p.Skip != 0 {
		skip = uint64(p.Skip)
	}
	if p.Take != 0 {
		take = uint64(p.Take)
	}
	return skip, take
}

func (c *queryResolver) Account(ctx context.Context, id string) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	r, err := c.server.accountClient.GetAccount(ctx, id)
	if err != nil {
		log.Println("Error Getting Account at Graphql ", err)
		return nil, err
	}
	return &Account{
		ID:   r.ID,
		Name: r.Name,
	}, nil
}

func (c *queryResolver) OrdersByAccount(ctx context.Context, accountID string) ([]*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	orderList, err := c.server.orderClient.GetOrdersForAccount(ctx, accountID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var orders []*Order
	for _, o := range orderList {
		order := &Order{
			ID:         o.ID,
			CreatedAt:  o.CreatedAt,
			TotalPrice: o.TotalPrice,
			Products:   []*Product{},
		}
		//order.CreatedAt, err = o.CreatedAt.UnmarshalBinary()MarshalBinary()
		if err != nil {
			log.Println(err)
		}
		for _, p := range o.Products {
			order.Products = append(order.Products, &Product{
				ID:          p.ID,
				Description: p.Description,
				Price:       p.Price,
				Name:        p.Name,
			})
		}
		orders = append(orders, order)
	}

	return orders, nil
}
