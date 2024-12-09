package server

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
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

	//cartClient, err := cart.NewClient(cartUrl)
	//if err != nil {
	//	//accountClient.Close()
	//	return nil, err
	//}

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

func (s *Server) AddItem(ctx context.Context, cartName, accountId, guestId, productId string, quantity uint64) error {
	md := metadata.New(map[string]string{
		"UserID":  ctx.Value("UserID").(string),
		"CartID":  ctx.Value("CartID").(string),
		"GuestID": ctx.Value("GuestID").(string),
	})

	ctx = metadata.NewOutgoingContext(context.Background(), md)

	_, err := s.CartClient.AddItemToCart(ctx, &pb.AddToCartRequest{
		ProductId: productId,
		Quantity:  quantity,
	})
	return err
}

func (s *Server) ValidateGuestCartToken(ctx context.Context, token string) (string, error) {
	md := metadata.New(map[string]string{
		"Token": token,
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	res, err := s.CartClient.ValidateGuestCartToken(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return res.GuestId, nil
}

func (s *Server) IssueGuestCartToken(ctx context.Context) (string, error) {
	res, err := s.CartClient.IssueGuestCartToken(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return res.GuestToken, err
}
