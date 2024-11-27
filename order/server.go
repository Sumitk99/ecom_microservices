package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/account"
	"github.com/Sumitk99/ecom_microservices/catalog"
	"github.com/Sumitk99/ecom_microservices/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
	pb.UnimplementedOrderServiceServer
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return errors.New("Cannot connect to account microservice")
	}
	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return errors.New("Cannot connect to catalog microservice")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return errors.New("error initiating listener")
	}
	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, &grpcServer{
		service:                         s,
		accountClient:                   accountClient,
		catalogClient:                   catalogClient,
		UnimplementedOrderServiceServer: pb.UnimplementedOrderServiceServer{},
	})
	reflection.Register(server) // to avoid sharing .proto files to client
	err = server.Serve(lis)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (srv *grpcServer) PostOrder(ctx context.Context, req *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	_, err := srv.accountClient.GetAccount(ctx, req.AccountId)
	if err != nil {
		log.Println("Error getting account", err)
		return nil, errors.New("account not found")
	}

	productIDs := []string{}
	IdToQuantity := make(map[string]int)
	for _, p := range req.Products {
		productIDs = append(productIDs, p.ProductId)
		IdToQuantity[p.ProductId] = int(p.Quantity)
	}
	orderedProducts, err := srv.catalogClient.GetProducts(ctx, 0, 0, productIDs, "")
	if err != nil {
		log.Println("Error getting products", err)
		return nil, errors.New("product not found")
	}

	products := []OrderedProduct{}
	for _, p := range orderedProducts {
		products = append(products, OrderedProduct{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Quantity:    uint32(IdToQuantity[p.ID]),
		})
	}

	order, err := srv.service.PostOrder(ctx, req.AccountId, products)
	if err != nil {
		log.Println("Error posting order", err)
		return nil, errors.New("order not posted")
	}

	orderProto := &pb.Order{
		Id:         order.ID,
		AccountId:  order.AccountID,
		TotalPrice: order.TotalPrice,
		Products:   []*pb.Order_OrderProduct{},
	}
	orderProto.CreatedAt, err = order.CreatedAt.MarshalBinary()
	if err != nil {
		log.Println(err)
	}
	for _, p := range order.Products {
		orderProto.Products = append(orderProto.Products, &pb.Order_OrderProduct{
			Id:          p.ID,
			Description: p.Description,
			Price:       p.Price,
			Name:        p.Name,
			Quantity:    p.Quantity,
		})
	}

	return &pb.PostOrderResponse{
		Order: orderProto,
	}, nil
}

func (srv *grpcServer) GetOrdersForAccount(ctx context.Context, req *pb.GetOrdersForAccountRequest) (*pb.GetOrdersForAccountResponse, error) {
	accountOrders, err := srv.service.GetOrdersForAccount(ctx, req.AccountId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	orders := []*pb.Order{}
	for _, o := range accountOrders {
		orderProto := &pb.Order{
			Id:         o.ID,
			AccountId:  o.AccountID,
			TotalPrice: o.TotalPrice,
			Products:   []*pb.Order_OrderProduct{},
		}
		orderProto.CreatedAt, err = o.CreatedAt.MarshalBinary()
		if err != nil {
			log.Println(err)
		}
		for _, p := range o.Products {
			orderProto.Products = append(orderProto.Products, &pb.Order_OrderProduct{
				Id:          p.ID,
				Description: p.Description,
				Price:       p.Price,
				Name:        p.Name,
				Quantity:    p.Quantity,
			})
		}
		orders = append(orders, orderProto)
	}
	return &pb.GetOrdersForAccountResponse{
		Orders: orders,
	}, nil
}
