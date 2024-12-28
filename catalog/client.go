package catalog

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/catalog/models"
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

func (c *Client) PostProduct(ctx context.Context, name, description string, price float64) (*models.Product, error) {
	res, err := c.Service.PostProduct(ctx, &pb.PostProductRequest{
		Title:       name,
		Description: description,
		Price:       price,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	colors := []models.Color{}
	for _, color := range res.Product.Colors {
		colors = append(colors, models.Color{
			ColorName: color.ColorName,
			Hex:       color.Hex,
		})
	}
	return &models.Product{
		ID:          res.Product.ProductId,
		Name:        res.Product.Title,
		Description: res.Product.Description,
		Price:       res.Product.Price,
		Category:    res.Product.Category,
		Stock:       res.Product.AvailableQuantity,
		Locations:   res.Product.Locations,
		ImageUrl:    res.Product.ImageURL,
		Sizes:       res.Product.Sizes,
		Colors:      colors,
		SellerID:    res.Product.SellerId,
		SellerName:  res.Product.SellerName,
	}, nil
}

func (c *Client) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	res, err := c.Service.GetProduct(
		ctx,
		&pb.GetProductRequest{Id: id},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	colors := []models.Color{}
	for _, color := range res.Product.Colors {
		colors = append(colors, models.Color{
			ColorName: color.ColorName,
			Hex:       color.Hex,
		})
	}
	return &models.Product{
		ID:          res.Product.ProductId,
		Name:        res.Product.Title,
		Description: res.Product.Description,
		Price:       res.Product.Price,
		Category:    res.Product.Category,
		Stock:       res.Product.AvailableQuantity,
		Locations:   res.Product.Locations,
		ImageUrl:    res.Product.ImageURL,
		Sizes:       res.Product.Sizes,
		Colors:      colors,
		SellerID:    res.Product.SellerId,
		SellerName:  res.Product.SellerName,
	}, nil
}

func (c *Client) GetProducts(ctx context.Context, skip, take uint64, ids []string, query string) ([]models.Product, error) {
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
	products := []models.Product{}
	for _, p := range res.Products {
		products = append(products, models.Product{
			ID:         p.ProductId,
			Name:       p.Title,
			Price:      p.Price,
			SellerName: p.SellerName,
			ImageUrl:   p.ImageURL,
		})
	}
	return products, nil
}
