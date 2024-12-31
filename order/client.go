package order

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/order/models"
	"github.com/Sumitk99/ecom_microservices/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Conn    *grpc.ClientConn
	Service pb.OrderServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(
		url, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}
	service := pb.NewOrderServiceClient(conn)
	return &Client{Conn: conn, Service: service}, nil
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) PostOrder(ctx context.Context, accountID string, products []models.OrderedProduct) (*models.Order, error) {
	protoProducts := []*pb.PostOrderRequest_OrderProduct{}
	for _, p := range products {
		protoProducts = append(protoProducts, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity:  p.Quantity,
		})
	}
	res, err := c.Service.PostOrder(ctx, &pb.PostOrderRequest{
		AccountId: accountID,
		Products:  protoProducts,
	})

	if err != nil {
		return nil, err
	}
	newOrder := res.Order
	newOrderCreatedAt := res.Order.CreatedAt
	return &models.Order{
		ID:         newOrder.OrderId,
		CreatedAt:  newOrderCreatedAt,
		TotalPrice: newOrder.TotalPrice,
		AccountID:  accountID,
		Products:   products,
	}, nil
}

//func (c *Client) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
//	res, err := c.Service.GetOrdersForAccount(
//		ctx,
//		&pb.GetOrdersForAccountRequest{AccountId: accountID})
//
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	orders := []Order{}
//
//	for _, o := range res.Orders {
//		newOrder := Order{
//			ID:         o.Id,
//			TotalPrice: o.TotalPrice,
//			AccountID:  accountID,
//			Products:   []OrderedProduct{},
//		}
//		newOrder.CreatedAt = o.CreatedAt
//		for _, p := range o.Products {
//			newOrder.Products = append(newOrder.Products, OrderedProduct{
//				ID:       p.Id,
//				Name:     p.Name,
//				Price:    p.Price,
//				Quantity: p.Quantity,
//			})
//		}
//		orders = append(orders, newOrder)
//	}
//
//	return orders, nil
//}
