package server

import (
	"context"
	"errors"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) SignUp(ctx context.Context, name, password, email, phone, user_type string) (*pb.SignUpResponse, error) {
	res, err := s.AccountClient.SignUp(ctx, &pb.SignUpRequest{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
		UserType: user_type,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server) Login(ctx context.Context, email, phone, password string) (*pb.LoginResponse, error) {
	if len(email) == 0 && len(phone) == 0 {
		return nil, errors.New("email or phone is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}
	LoginPayload := &pb.LoginRequest{
		Password: password,
	}
	if email != "" {
		LoginPayload.ContactMethod = &pb.LoginRequest_Email{Email: email}
	} else {
		LoginPayload.ContactMethod = &pb.LoginRequest_Phone{Phone: phone}
	}

	res, err := s.AccountClient.Login(ctx, LoginPayload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, err
}

func (s *Server) GetAccount(ctx context.Context) (*pb.AccountResponse, error) {
	md := metadata.New(map[string]string{
		"UserID": ctx.Value("UserID").(string),
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	res, err := s.AccountClient.GetAccount(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server) GetAccounts(ctx context.Context, skip, take uint64) (*pb.GetAccountsResponse, error) {
	res, err := s.AccountClient.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Skip: skip,
			Take: take,
		})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server) Authentication(ctx context.Context) (*pb.AccountResponse, error) {
	md := metadata.New(map[string]string{
		"authorization": ctx.Value("authorization").(string),
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	res, err := s.AccountClient.Authentication(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return res, err
}
