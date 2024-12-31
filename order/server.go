package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/account"
	"github.com/Sumitk99/ecom_microservices/catalog"
	"github.com/Sumitk99/ecom_microservices/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
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

	lis, err := net.Listen("tcp", fmt.Sprintf(fmt.Sprintf(":%d", port)))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return errors.New(fmt.Sprintf("Cannot listen %s", err))
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
	//_, err := srv.accountClient.GetAccount(ctx)
	//if err != nil {
	//	log.Println("Error getting account", err)
	//	return nil, errors.New("account not found")
	//}
	if req.MethodOfPayment != "COD" {
		return &pb.PostOrderResponse{
			Message: "Other Payment to be Available Soon. Stay Tuned",
		}, nil
	}
	if len(req.MethodOfPayment) == 0 {
		return nil, errors.New("Select A Payment Method to Continue")
	}
	if req.MethodOfPayment != "COD" && len(req.TransactionId) == 0 {
		return nil, errors.New("No Transaction ID Found")
	}

	log.Println("started posting order")
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
			ImageURL:    p.ImageUrl,
		})
	}

	order, err := srv.service.PostOrder(ctx, req.AccountId, req.MethodOfPayment, req.TransactionId, req.PaymentStatus, products)
	if err != nil {
		log.Println("Error posting order", err)
		return nil, errors.New("order not posted")
	}

	orderProto := &pb.Order{
		Id:              order.ID,
		AccountId:       order.AccountID,
		TotalPrice:      order.TotalPrice,
		MethodOfPayment: order.MethodOfPayment,
		TransactionId:   order.TransactionID,
		Products:        []*pb.Order_OrderProduct{},
	}

	orderProto.CreatedAt = order.CreatedAt
	for _, p := range order.Products {
		orderProto.Products = append(orderProto.Products, &pb.Order_OrderProduct{
			Id:       p.ID,
			Price:    p.Price,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}
	orderProto.ETA = time.Now().Add(time.Hour * 7 * 24).String()
	if req.MethodOfPayment == "COD" {
		orderProto.PaymentStatus = "Cash On Delivery"
		orderProto.OrderStatus = "Order Placed"
	}
	return &pb.PostOrderResponse{
		Order:   orderProto,
		Message: "Order Successfully Placed",
	}, nil
}

func (srv *grpcServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}
	acc := md.Get("UserID")
	if len(acc) == 0 || len(acc[0]) == 0 {
		return nil, errors.New("not Enough Data to Get Order, Login to Access your Orders")
	}

	OrderID := req.OrderId
	accountID := acc[0]
	order, err := srv.service.GetOrder(ctx, OrderID, accountID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	orderRes := new(pb.Order)

	orderRes = &pb.Order{
		Id:              order.ID,
		AccountId:       order.AccountID,
		TotalPrice:      order.TotalPrice,
		MethodOfPayment: order.MethodOfPayment,
		TransactionId:   order.TransactionID,
		CreatedAt:       order.CreatedAt,
		PaymentStatus:   order.PaymentStatus,
		ETA:             order.ETA,
		OrderStatus:     order.OrderStatus,
	}
	for _, p := range order.Products {
		orderRes.Products = append(orderRes.Products, &pb.Order_OrderProduct{
			Id:       p.ID,
			Price:    p.Price,
			Name:     p.Name,
			Quantity: p.Quantity,
			ImageURL: p.ImageURL,
		})
		fmt.Printf("imageURL: %s\n", p.ImageURL)
	}
	return &pb.GetOrderResponse{
		Order: orderRes,
	}, nil
}

func (srv *grpcServer) GetOrdersForAccount(ctx context.Context, req *pb.GetOrdersForAccountRequest) (*pb.GetOrdersForAccountResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}
	acc := md.Get("UserID")
	if len(acc) == 0 || len(acc[0]) == 0 {
		return nil, errors.New("not Enough Data to Get Order, Login to Access your Orders")
	}
	log.Println(acc[0], req.AccountId)
	if acc[0] != req.AccountId {
		return nil, errors.New("Unauthorized to Access This Resource")
	}
	accountOrders, err := srv.service.GetOrdersForAccount(ctx, req.AccountId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	Orders := &pb.GetOrdersForAccountResponse{
		Orders: []*pb.GetOrdersForAccountResponse_Order{},
	}

	for _, order := range accountOrders {
		Orders.Orders = append(Orders.Orders, &pb.GetOrdersForAccountResponse_Order{
			OrderId:     order.OrderId,
			CreatedAt:   order.CreatedAt,
			TotalPrice:  order.TotalPrice,
			ETA:         order.ETA,
			OrderStatus: order.OrderStatus,
		})
	}
	return Orders, err
}
