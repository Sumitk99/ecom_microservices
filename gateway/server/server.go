package server

import (
	"github.com/Sumitk99/ecom_microservices/account"
)

type Server struct {
	AccountClient *account.Client
	//catalogClient *catalog.Client
	//orderClient   *order.Client
}

func NewGinServer(accountUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}
	//catalogClient, err := catalog.NewClient(catalogUrl)
	//if err != nil {
	//	accountClient.Close()
	//	return nil, err
	//}
	//orderClient, err := order.NewClient(orderUrl)
	//if err != nil {
	//	accountClient.Close()
	//	catalogClient.Close()
	//	return nil, err
	//}
	return &Server{accountClient}, nil
}
