package main

import (
	"context"
	"log"
	"time"
)

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	orderList, err := r.server.orderClient.GetOrdersForAccount(ctx, obj.ID)
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
