package catalog

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/catalog/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	Conn    *grpc.ClientConn
	Service pb.CatalogServiceClient
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
	service := pb.NewCatalogServiceClient(conn)
	return &Client{Conn: conn, Service: service}, nil
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) PostProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	res, err := c.Service.PostProduct(ctx, &pb.PostProductRequest{
		Name:        name,
		Description: description,
		Price:       price,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Product{
		ID:          res.Product.Id,
		Name:        res.Product.Name,
		Description: res.Product.Description,
		Price:       res.Product.Price,
	}, nil
}

func (c *Client) GetProduct(ctx context.Context, id string) (*Product, error) {
	res, err := c.Service.GetProduct(
		ctx,
		&pb.GetProductRequest{Id: id},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Product{
		ID:          res.Product.Id,
		Name:        res.Product.Name,
		Description: res.Product.Description,
		Price:       res.Product.Price,
	}, nil
}

func (c *Client) GetProducts(ctx context.Context, skip, take uint64, ids []string, query string) ([]Product, error) {
	res, err := c.Service.GetProducts(
		ctx,
		&pb.GetProductsRequest{
			Skip:  skip,
			Take:  take,
			Ids:   ids,
			Query: query,
		},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []Product{}
	for _, p := range res.Products {
		products = append(products, Product{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return products, nil
}
