package account

// go generate protoc --go_out=plugins=grpc:./pb account.proto
import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/account/models"
	"github.com/Sumitk99/ecom_microservices/account/pb"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type grpcServer struct {
	service Service
	pb.UnimplementedAccountServiceServer
}

func ListenGRPC(s Service, port string) error {

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
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

func (s *grpcServer) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.Address, error) {
	add := models.AddAddressRequest{
		Name:          req.Name,
		Phone:         req.Phone,
		Street:        req.Street,
		City:          req.City,
		State:         req.State,
		ZipCode:       req.ZipCode,
		Country:       req.Country,
		IsDefault:     req.IsDefault,
		ApartmentUnit: req.ApartmentUnit,
	}
	validationErr := validator.New().Struct(add)
	if validationErr != nil {
		log.Println(validationErr)
		return nil, validationErr
	}

	res, err := s.service.AddAddress(ctx, add)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(fmt.Sprintf("Error adding address: %s", err.Error()))
	}
	return &pb.Address{
		Name:          res.Name,
		Phone:         res.Phone,
		Street:        res.Street,
		City:          res.City,
		State:         res.State,
		ZipCode:       res.ZipCode,
		Country:       res.Country,
		IsDefault:     res.IsDefault,
		ApartmentUnit: res.ApartmentUnit,
		AddressId:     res.AddressID,
		UserId:        res.UserID,
		CreatedAt:     res.CreatedAt,
	}, nil
}

func (s *grpcServer) GetAddresses(ctx context.Context, req *emptypb.Empty) (*pb.Addresses, error) {
	add, err := s.service.GetAddresses(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Error getting addresses: %s", err))
	}
	addresses := []*pb.Address{}
	for _, a := range add {
		addresses = append(addresses, &pb.Address{
			Name:          a.Name,
			Phone:         a.Phone,
			Street:        a.Street,
			City:          a.City,
			State:         a.State,
			ZipCode:       a.ZipCode,
			Country:       a.Country,
			IsDefault:     a.IsDefault,
			ApartmentUnit: a.ApartmentUnit,
			AddressId:     a.AddressID,
			UserId:        a.UserID,
			CreatedAt:     a.CreatedAt,
		})
	}
	return &pb.Addresses{
		Addresses: addresses,
	}, nil
}

func (s *grpcServer) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*emptypb.Empty, error) {
	err := s.service.DeleteAddress(ctx, req.AddressId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Error deleting address: %s", err))
	}
	return &emptypb.Empty{}, nil
}

func (s *grpcServer) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.Address, error) {
	add, err := s.service.GetAddress(ctx, req.AddressId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Error getting address: %s", err))
	}
	return &pb.Address{
		Name:          add.Name,
		Phone:         add.Phone,
		Street:        add.Street,
		City:          add.City,
		State:         add.State,
		ZipCode:       add.ZipCode,
		Country:       add.Country,
		IsDefault:     add.IsDefault,
		ApartmentUnit: add.ApartmentUnit,
		AddressId:     add.AddressID,
		UserId:        add.UserID,
		CreatedAt:     add.CreatedAt,
	}, nil
}
