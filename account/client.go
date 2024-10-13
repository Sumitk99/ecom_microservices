package account

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func (c *Client) PostAccount(ctx context.Context, name string) (*Account, error) {
	r, err := c.Service.PostAccount(ctx, &pb.PostAccountRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	res, err := c.Service.GetAccount(ctx, &pb.GetAccountRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:   res.Account.Id,
		Name: res.Account.Name,
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
