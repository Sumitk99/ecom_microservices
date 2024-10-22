package catalog

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/catalog/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func ListenGRPC(s Service, port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterCatalogServiceServer(srv, &grpcServer{
		service:                           s,
		UnimplementedCatalogServiceServer: pb.UnimplementedCatalogServiceServer{},
	})
	//pb.RegisterAccountServiceServer(srv, &grpcServer{ s})

	//reflection.Register(srv)
	err = srv.Serve(lis)
	return err
}

type grpcServer struct {
	pb.UnimplementedCatalogServiceServer
	service Service
}

func (s *grpcServer) PostProduct(ctx context.Context, r *pb.PostProductRequest) (*pb.PostProductResponse, error) {
	acc, err := s.service.PostProduct(ctx, r.Name, r.Description, r.Price)
	if err != nil {
		return nil, err
	}
	return &pb.PostProductResponse{Product: &pb.Product{
		Id:          acc.ID,
		Name:        acc.Name,
		Description: acc.Description,
		Price:       acc.Price,
	},
		Message: "Product Successfully Created",
	}, nil
}

func (s *grpcServer) GetProduct(ctx context.Context, r *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	acc, err := s.service.GetProduct(ctx, r.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.GetProductResponse{Product: &pb.Product{
		Id:          acc.ID,
		Name:        acc.Name,
		Description: acc.Description,
		Price:       acc.Price,
	},
	}, nil
}

func (s *grpcServer) GetProducts(ctx context.Context, r *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	var res []Product
	var err error

	if r.Query != "" {
		res, err = s.service.SearchProducts(ctx, r.Query, r.Skip, r.Take)
	} else if len(r.Ids) > 0 {
		res, err = s.service.GetProductByIDs(ctx, r.Ids)
	} else {
		res, err = s.service.GetProducts(ctx, r.Skip, r.Take)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []*pb.Product{}
	for _, p := range res {
		products = append(products, &pb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return &pb.GetProductsResponse{Products: products}, nil
}
