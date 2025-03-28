package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/cart/helper"
	"github.com/Sumitk99/ecom_microservices/cart/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const AlreadyExists = "Item already exists in selected cart"
const NoUserData = "not enough user Data provided"

type grpcServer struct {
	service       CartService
	catalogClient pb.CatalogServiceClient
	OrderClient   pb.OrderServiceClient
	pb.UnimplementedCartServiceServer
}

func ListenGRPC(s CartService, catalogURL, orderURL, port string) error {
	CatalogConn, err := grpc.Dial(catalogURL, grpc.WithInsecure())
	if err != nil {
		return errors.New("Cannot connect to order microservice")
	}

	CatalogClient := pb.NewCatalogServiceClient(CatalogConn)

	orderConn, err := grpc.Dial(orderURL, grpc.WithInsecure())
	if err != nil {
		return errors.New("Cannot connect to order microservice")
	}
	orderClient := pb.NewOrderServiceClient(orderConn)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterCartServiceServer(srv, &grpcServer{
		service:                        s,
		OrderClient:                    orderClient,
		catalogClient:                  CatalogClient,
		UnimplementedCartServiceServer: pb.UnimplementedCartServiceServer{},
	})
	reflection.Register(srv)
	err = srv.Serve(lis)
	return err
}

func (s *grpcServer) AddItemToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.CartResponse, error) {
	log.Println(fmt.Sprintf("AddItemToCart: %s %d", req.ProductId, req.Quantity))
	_, err := s.catalogClient.GetProduct(ctx, &pb.GetProductRequest{Id: req.ProductId})
	log.Println(fmt.Sprintf("fetched"))
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("product Not Found in catalog: %s", err))
	}
	log.Println(fmt.Sprintf("no error in fetching"))

	err = s.service.AddItem(ctx, req.ProductId, req.Quantity)
	if err != nil {
		if err.Error() == AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, AlreadyExists)
		}
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Error %s", err))
	}
	cart, err := s.GetCart(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cart, err
}

func (s *grpcServer) GetCart(ctx context.Context, req *emptypb.Empty) (*pb.CartResponse, error) {
	log.Println("GetCart")
	CartProducts, err := s.service.GetCartItems(ctx)
	if len(CartProducts) == 0 {
		return nil, errors.New("Cart is Empty")
	}
	IdToQuantity := make(map[string]uint64)

	productIds := helper.MakeProductArray(CartProducts, &IdToQuantity)
	//products, err := s.catalogClient.GetProducts(ctx, 0, 0, *productIds, "")
	products, err := s.catalogClient.GetProducts(ctx, &pb.GetProductsRequest{
		Ids: *productIds,
	})

	if err != nil {
		return nil, err
	}

	CartItems, totalPrice := helper.ProcessCart(products.Products, IdToQuantity)

	md, _ := metadata.FromIncomingContext(ctx)
	cart, guestId := md.Get("CartID"), md.Get("GuestID")
	var IdMethod string
	if len(cart) > 0 {
		IdMethod = cart[0]
	} else {
		IdMethod = guestId[0]
	}
	return &pb.CartResponse{
		Cart: &pb.Cart{
			CartId:     IdMethod,
			Items:      CartItems,
			TotalPrice: totalPrice,
		},
	}, err
}

func (s *grpcServer) RemoveItemFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.CartResponse, error) {
	//_, err := s.catalogClient.GetProduct(ctx, &pb.GetProductRequest{Id: req.ProductId})
	//if err != nil {
	//	log.Println(err)
	//	return nil, errors.New(fmt.Sprintf("product Not Found in catalog: %s", err))
	//}

	err := s.service.DeleteItem(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot delete item from cart : %s\n", err))
	}
	cart, err := s.GetCart(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cart, err
}

func (s *grpcServer) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) (*pb.CartResponse, error) {
	_, err := s.catalogClient.GetProduct(ctx, &pb.GetProductRequest{Id: req.ProductId})
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("product Not Found in catalog: %s", err))
	}

	err = s.service.UpdateItem(ctx, req.ProductId, req.UpdatedQuantity)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot update item in cart : %s\n", err))
	}
	cart, err := s.GetCart(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cart, err
}

func (s *grpcServer) DeleteCart(ctx context.Context, req *emptypb.Empty) (*pb.CartResponse, error) {
	err := s.service.DeleteCart(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot delete cart : %s\n", err))
	}
	return &pb.CartResponse{}, err
}

func (s *grpcServer) IssueGuestCartToken(ctx context.Context, req *emptypb.Empty) (*pb.IssueGuestCartTokenResponse, error) {
	guestToken, err := s.service.IssueGuestToken(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot issue guest id : %s\n", err))
	}
	return &pb.IssueGuestCartTokenResponse{
		GuestToken: guestToken,
	}, err
}

func (s *grpcServer) ValidateGuestCartToken(ctx context.Context, req *emptypb.Empty) (*pb.ValidateGuestCartTokenResponse, error) {
	guestId, err := s.service.ValidateGuestId(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot validate guest id : %s\n", err))
	}
	return &pb.ValidateGuestCartTokenResponse{
		GuestId: guestId,
	}, err
}

func (s *grpcServer) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.PostOrderResponse, error) {
	if len(req.MethodOfPayment) == 0 {
		return nil, errors.New("Select A Payment Method to Continue")
	}
	if req.MethodOfPayment != "COD" && len(req.TransactionId) == 0 {
		return nil, errors.New("No Transaction ID Found")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}
	account, cart := md.Get("UserID"), md.Get("CartID")
	var products *pb.CartResponse
	var err error

	if len(account) > 0 && len(cart) > 0 && len(account[0]) > 0 && len(cart[0]) > 0 {
		products, err = s.GetCart(ctx, &emptypb.Empty{})
		if err != nil {
			log.Println(fmt.Sprintf("error getting cart : %s\n", err))
			return nil, errors.New(fmt.Sprintf("cannot checkout cart : %s\n", err))
		}
	} else {
		return nil, errors.New("not Enough Data to Checkout Cart")
	}
	log.Println("Got Cart ITEMs")

	orderReq := new(pb.PostOrderRequest)
	orderReq.AccountId, orderReq.MethodOfPayment = account[0], req.MethodOfPayment
	orderReq.TransactionId = req.TransactionId
	orderReq.AddressId = req.AddressId

	if req.MethodOfPayment == "COD" {
		orderReq.PaymentStatus = "COD"
	}

	for _, p := range products.Cart.Items {
		orderReq.Products = append(orderReq.Products, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ProductId,
			Quantity:  uint32(p.Quantity),
		})
	}
	// To check if the all the products in the cart is available in the catalog
	//products, err := s.catalogClient.GetProducts(ctx, 0, 0, nil, "")
	res, err := s.OrderClient.PostOrder(ctx, orderReq)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot checkout cart : %s\n", err))
	}
	//_, err = s.DeleteCart(ctx, &emptypb.Empty{})
	return res, err
}
