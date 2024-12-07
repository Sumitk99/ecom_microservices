package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/cart/helper"
	"github.com/Sumitk99/ecom_microservices/cart/pb"
	"github.com/Sumitk99/ecom_microservices/catalog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	service       CartService
	catalogClient *catalog.Client
	pb.UnimplementedCartServiceServer
}

func ListenGRPC(s CartService, catalogURL, port string) error {
	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		return errors.New("Cannot connect to catalog microservice")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterCartServiceServer(srv, &grpcServer{
		service:                        s,
		catalogClient:                  catalogClient,
		UnimplementedCartServiceServer: pb.UnimplementedCartServiceServer{},
	})
	reflection.Register(srv)
	err = srv.Serve(lis)
	return err
}

func (s *grpcServer) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	_, err := s.catalogClient.GetProduct(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("product Not Found in catalog: %s", err))
	}

	err = s.service.AddItem(ctx, req.ProductId, int(req.Quantity))
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot add item to cart : %s\n", err))
	}

	ResponseCart, err := s.GetCart(ctx, &pb.GetCartRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.AddItemResponse{
		Cart: ResponseCart.Cart,
	}, err
}

func (s *grpcServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	CartProducts, err := s.service.GetCartItems(ctx)
	if len(CartProducts) == 0 {
		return nil, errors.New("Cart is Empty")
	}
	IdToQuantity := make(map[string]uint64)

	productIds := helper.MakeProductArray(CartProducts, &IdToQuantity)
	products, err := s.catalogClient.GetProducts(ctx, 0, 0, *productIds, "")
	if err != nil {
		return nil, err
	}

	CartItems, totalPrice := helper.ProcessCart(products, IdToQuantity)

	md, _ := metadata.FromIncomingContext(ctx)
	cart, guestId := md.Get("CartID"), md.Get("GuestID")
	var IdMethod string
	if len(cart) > 0 {
		IdMethod = cart[0]
	} else {
		IdMethod = guestId[0]
	}
	return &pb.GetCartResponse{
		Cart: &pb.Cart{
			CartId:     IdMethod,
			Items:      CartItems,
			TotalPrice: totalPrice,
		},
	}, err
}

func (s *grpcServer) RemoveItem(ctx context.Context, req *pb.RemoveItemRequest) (*pb.RemoveItemResponse, error) {
	_, err := s.catalogClient.GetProduct(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("product Not Found in catalog: %s", err))
	}

	err = s.service.DeleteItem(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot delete item from cart : %s\n", err))
	}
	ResponseCart, err := s.GetCart(ctx, &pb.GetCartRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.RemoveItemResponse{
		Cart: ResponseCart.Cart,
	}, err

}

func (s *grpcServer) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error) {
	_, err := s.catalogClient.GetProduct(ctx, req.ProductId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("product Not Found in catalog: %s", err))
	}

	err = s.service.UpdateItem(ctx, req.ProductId, req.UpdatedQuantity)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot update item in cart : %s\n", err))
	}
	ResponseCart, err := s.GetCart(ctx, &pb.GetCartRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.UpdateCartResponse{
		Cart: ResponseCart.Cart,
	}, err
}

func (s *grpcServer) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error) {
	err := s.service.DeleteCart(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("cannot delete cart : %s\n", err))
	}
	return &pb.DeleteCartResponse{}, err
}
