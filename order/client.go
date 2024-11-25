package order

import "google.golang.org/grpc"

type Client struct {
	client *grpc.ClientConn
	Service pb.
}