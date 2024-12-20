package server

import (
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	AccountClient pb.AccountServiceClient
	//catalogClient *catalog.Client
	//orderClient   *order.Client
	CartClient pb.CartServiceClient
}

func NewGinServer(accountUrl, cartUrl string) (*Server, error) {
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
		return nil, err
	}
	AccountService := pb.NewAccountServiceClient(AccountConn)

	//catalogClient, err := catalog.NewClient(catalogUrl)
	//if err != nil {
	//	accountClient.Close()
	//	return nil, err
	//}
	return &Server{
		AccountClient: AccountService,
		CartClient:    CartService,
	}, nil
}
