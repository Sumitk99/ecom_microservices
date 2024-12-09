package account

// go generate protoc --go_out=plugins=grpc:./pb account.proto
import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/account/pb"
	"github.com/Sumitk99/ecom_microservices/catalog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type grpcServer struct {
	service       Service
	catalogClient *catalog.Client
	pb.UnimplementedAccountServiceServer
}

func ListenGRPC(s Service, port string) error {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv, &grpcServer{
		service:                           s,
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
	})
	reflection.Register(srv)
	err = srv.Serve(lis)
	return err
}

func (s *grpcServer) SignUp(ctx context.Context, r *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	acc, err := s.service.SignUp(ctx, r.Name, r.Password, r.Email, r.Phone, r.UserType)
	if err != nil {
		return nil, err
	}
	return &pb.SignUpResponse{Account: &pb.Account{
		Id:       acc.ID,
		Name:     acc.Name,
		Email:    acc.Email,
		Phone:    acc.Phone,
		UserType: acc.UserType,
	},
		Message: "Account Successfully Created",
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *emptypb.Empty) (*pb.AccountResponse, error) {
	log.Println("Server Side: %s", ctx.Value("UserID"))
	acc, err := s.service.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	log.Println(acc)
	return &pb.AccountResponse{
		Account: &pb.Account{
			Id:       acc.ID,
			Name:     acc.Name,
			Email:    acc.Email,
			Phone:    acc.Phone,
			UserType: acc.UserType,
		},
	}, nil
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	acc, err := s.service.Login(ctx, req.GetEmail(), req.GetPhone(), req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Account: &pb.Account{
			Id:       acc.ID,
			Name:     acc.Name,
			Email:    acc.Email,
			Phone:    acc.Phone,
			UserType: acc.UserType,
		},
		JWT_Token:     acc.Token,
		Refresh_Token: acc.RefreshToken,
	}, nil
}

func (s *grpcServer) Authentication(ctx context.Context, r *emptypb.Empty) (*pb.AccountResponse, error) {
	acc, err := s.service.Authentication(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.AccountResponse{
		Account: &pb.Account{
			Id:       acc.ID,
			Name:     acc.Name,
			Email:    acc.Email,
			Phone:    acc.Phone,
			UserType: acc.UserType,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	if ctx.Value("USER_TYPE") != "ADMIN" {
		log.Printf("%s is Unauthorized to access this resource\n", ctx.Value("USER_TYPE"))
		return nil, errors.New("Unauthorized to access this resource")
	}
	acc, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, a := range acc {
		accounts = append(accounts, &pb.Account{
			Id: a.ID, Name: a.Name,
		})
	}
	return &pb.GetAccountsResponse{Accounts: accounts}, nil
}
