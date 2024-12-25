package server

import (
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	AccountClient pb.AccountServiceClient
	//catalogClient *catalog.Client
	OrderClient pb.OrderServiceClient
	CartClient  pb.CartServiceClient
}

func NewGinServer(accountUrl, cartUrl, orderUrl string) (*Server, error) {
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

	//catalogClient, err := catalog.NewClient(catalogUrl)
	//if err != nil {
	//	accountClient.Close()
	//	return nil, err
	//}
	return &Server{
		AccountClient: AccountService,
		CartClient:    CartService,
		OrderClient:   OrderService,
	}, nil
}
