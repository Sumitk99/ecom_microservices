package account

import (
	"context"
	"errors"
	"github.com/Sumitk99/ecom_microservices/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

type Client struct {
	Conn    *grpc.ClientConn
	Service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(
		url, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}
	service := pb.NewAccountServiceClient(conn)
	return &Client{Conn: conn, Service: service}, nil
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) SignUp(ctx context.Context, name, password, email, phone, user_type string) (*Account, error) {
	r, err := c.Service.SignUp(ctx, &pb.SignUpRequest{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
		UserType: user_type,
	})
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:       r.Account.Id,
		Name:     r.Account.Name,
		Email:    r.Account.Email,
		Phone:    r.Account.Phone,
		UserType: r.Account.UserType,
	}, nil
}

func (c *Client) Login(ctx context.Context, email, phone, password string) (*Account, string, string, error) {
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

	res, err := c.Service.Login(ctx, LoginPayload)
	if err != nil {
		log.Println(err)
		return nil, "", "", err
	}
	return &Account{
		ID:       res.Account.Id,
		Name:     res.Account.Name,
		Email:    res.Account.Email,
		Phone:    res.Account.Phone,
		UserType: res.Account.UserType,
	}, res.JWT_Token, res.Refresh_Token, err
}

func (c *Client) GetAccount(ctx context.Context) (*Account, error) {
	log.Println("Client Side: %s", ctx.Value("UserID"))
	md := metadata.New(map[string]string{
		"UserID": ctx.Value("UserID").(string),
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	res, err := c.Service.GetAccount(ctx, &pb.GetAccountRequest{})
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:       res.Account.Id,
		Name:     res.Account.Name,
		Email:    res.Account.Email,
		Phone:    res.Account.Phone,
		UserType: res.Account.UserType,
	}, nil
}

func (c *Client) GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {
	res, err := c.Service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Skip: skip,
			Take: take,
		})
	if err != nil {
		return nil, err
	}
	accounts := make([]Account, len(res.Accounts))
	for i, a := range res.Accounts {
		accounts[i] = Account{
			ID:   a.Id,
			Name: a.Name,
		}
	}
	return accounts, nil
}

func (c *Client) Authentication(ctx context.Context) (*Account, error) {
	md := metadata.New(map[string]string{
		"authorization": ctx.Value("authorization").(string),
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	res, err := c.Service.Authentication(ctx, &pb.AuthenticationRequest{})
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:       res.Account.Id,
		Name:     res.Account.Name,
		Email:    res.Account.Email,
		Phone:    res.Account.Phone,
		UserType: res.Account.UserType,
	}, nil
}
