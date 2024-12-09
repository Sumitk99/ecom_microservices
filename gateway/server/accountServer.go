package server

import (
	"context"
	"errors"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) SignUp(ctx context.Context, name, password, email, phone, user_type string) (*models.Account, error) {
	r, err := s.AccountClient.SignUp(ctx, &pb.SignUpRequest{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
		UserType: user_type,
	})
	if err != nil {
		return nil, err
	}
	return &models.Account{
		ID:       r.Account.Id,
		Name:     r.Account.Name,
		Email:    r.Account.Email,
		Phone:    r.Account.Phone,
		UserType: r.Account.UserType,
	}, nil
}

func (s *Server) Login(ctx context.Context, email, phone, password string) (*models.Account, string, string, error) {
	if email == "" && phone == "" {
		return nil, "", "", errors.New("email or phone is required")
	}
	if password == "" {
		return nil, "", "", errors.New("password is required")
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
		return nil, "", "", err
	}
	return &models.Account{
		ID:       res.Account.Id,
		Name:     res.Account.Name,
		Email:    res.Account.Email,
		Phone:    res.Account.Phone,
		UserType: res.Account.UserType,
	}, res.JWT_Token, res.Refresh_Token, err
}

func (s *Server) GetAccount(ctx context.Context) (*models.Account, error) {
	log.Println("Client Side: %s", ctx.Value("UserID"))
	md := metadata.New(map[string]string{
		"UserID": ctx.Value("UserID").(string),
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	res, err := s.AccountClient.GetAccount(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return &models.Account{
		ID:       res.Account.Id,
		Name:     res.Account.Name,
		Email:    res.Account.Email,
		Phone:    res.Account.Phone,
		UserType: res.Account.UserType,
	}, nil
}

func (s *Server) GetAccounts(ctx context.Context, skip, take uint64) ([]models.Account, error) {
	res, err := s.AccountClient.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Skip: skip,
			Take: take,
		})
	if err != nil {
		return nil, err
	}
	accounts := make([]models.Account, len(res.Accounts))
	for i, a := range res.Accounts {
		accounts[i] = models.Account{
			ID:   a.Id,
			Name: a.Name,
		}
	}
	return accounts, nil
}

func (s *Server) Authentication(ctx context.Context) (*models.Account, error) {
	md := metadata.New(map[string]string{
		"authorization": ctx.Value("authorization").(string),
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	res, err := s.AccountClient.Authentication(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return &models.Account{
		ID:       res.Account.Id,
		Name:     res.Account.Name,
		Email:    res.Account.Email,
		Phone:    res.Account.Phone,
		UserType: res.Account.UserType,
	}, nil
}
