package account

// go generate protoc --go_out=plugins=grpc:./pb account.proto
import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

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
	//pb.RegisterAccountServiceServer(srv, &grpcServer{ s})

	reflection.Register(srv)
	err = srv.Serve(lis)
	return err
}

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

//type grpcServer struct {
//	service Service
//}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	acc, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   acc.ID,
		Name: acc.Name,
	},
		Message: "Account Successfully Created",
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	acc, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   acc.ID,
			Name: acc.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
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
