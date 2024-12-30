package server

import (
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"github.com/cloudinary/cloudinary-go/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Server struct {
	AccountClient     pb.AccountServiceClient
	CatalogClient     pb.CatalogServiceClient
	OrderClient       pb.OrderServiceClient
	CartClient        pb.CartServiceClient
	CloudinaryStorage *cloudinary.Cloudinary
}

func NewGinServer(accountUrl, cartUrl, orderUrl, catalogUrl string) (*Server, error) {
	CartConn, err := grpc.NewClient(
		cartUrl, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}
	CartService := pb.NewCartServiceClient(CartConn)

	AccountConn, err := grpc.NewClient(
		accountUrl, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		CartConn.Close()
		return nil, err
	}
	AccountService := pb.NewAccountServiceClient(AccountConn)

	OrderConn, err := grpc.NewClient(
		orderUrl, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		CartConn.Close()
		AccountConn.Close()
		return nil, err
	}
	OrderService := pb.NewOrderServiceClient(OrderConn)

	CatalogConn, err := grpc.NewClient(
		catalogUrl, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		CartConn.Close()
		AccountConn.Close()
		OrderConn.Close()
		return nil, err
	}
	CatalogService := pb.NewCatalogServiceClient(CatalogConn)

	cld, err := cloudinary.NewFromParams("dwd3oedmz", "919171763669784", "sXGGnuvVdcHVG4VZZq7fJJCrxcA")
	if err != nil {
		log.Println("Failed to initialize Cloudinary: %v", err)
		CartConn.Close()
		AccountConn.Close()
		OrderConn.Close()
		CatalogConn.Close()
		return nil, err
	}

	return &Server{
		AccountClient:     AccountService,
		CartClient:        CartService,
		OrderClient:       OrderService,
		CatalogClient:     CatalogService,
		CloudinaryStorage: cld,
	}, nil
}
