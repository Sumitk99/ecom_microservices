package catalog

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/catalog/models"
	"github.com/Sumitk99/ecom_microservices/catalog/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func ListenGRPC(s Service, port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterCatalogServiceServer(srv, &grpcServer{
		service:                           s,
		UnimplementedCatalogServiceServer: pb.UnimplementedCatalogServiceServer{},
	})
	//pb.RegisterAccountServiceServer(srv, &grpcServer{ s})

	reflection.Register(srv)
	err = srv.Serve(lis)
	return err
}

type grpcServer struct {
	pb.UnimplementedCatalogServiceServer
	service Service
}

func (s *grpcServer) PostProduct(ctx context.Context, r *pb.PostProductRequest) (*pb.PostProductResponse, error) {
	colors := []models.Color{}
	for _, color := range r.Colors {
		colors = append(colors, models.Color{
			ColorName: color.ColorName,
			Hex:       color.Hex,
		})
	}
	prod, err := s.service.PostProduct(ctx, r.Title, r.Description, r.SellerId, r.SellerName, r.ImageURL, r.Category, r.Price, r.AvailableQuantity, r.Locations, r.Sizes, colors)
	if err != nil {
		return nil, err
	}
	return &pb.PostProductResponse{
		Product: &pb.Product{
			ProductId:         prod.ID,
			Title:             prod.Name,
			Description:       prod.Description,
			Price:             prod.Price,
			Category:          prod.Category,
			AvailableQuantity: prod.Stock,
			Locations:         prod.Locations,
			ImageURL:          prod.ImageUrl,
			Sizes:             prod.Sizes,
			Colors:            r.Colors,
			SellerId:          prod.SellerID,
			SellerName:        prod.SellerName,
		},
		Message: "Product Successfully Added",
	}, nil
}

func (s *grpcServer) GetProduct(ctx context.Context, r *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	res, err := s.service.GetProduct(ctx, r.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	colors := []*pb.Color{}
	for _, color := range res.Colors {
		colors = append(colors, &pb.Color{
			ColorName: color.ColorName,
			Hex:       color.Hex,
		})
	}
	return &pb.GetProductResponse{
		Product: &pb.Product{
			ProductId:         res.ID,
			Title:             res.Name,
			Description:       res.Description,
			Price:             res.Price,
			Category:          res.Category,
			AvailableQuantity: res.Stock,
			Locations:         res.Locations,
			ImageURL:          res.ImageUrl,
			Sizes:             res.Sizes,
			Colors:            colors,
			SellerId:          res.SellerID,
			SellerName:        res.SellerName,
		},
	}, nil
}

func (s *grpcServer) GetProducts(ctx context.Context, r *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	var res []models.Product
	var err error
	log.Println(r)
	if len(r.Query) > 0 {
		log.Println("Searching for products")
		res, err = s.service.SearchProducts(ctx, r.Query, r.Skip, r.Take)
	} else if len(r.Ids) > 0 {
		log.Println("Getting products by IDs")
		res, err = s.service.GetProductByIDs(ctx, r.Ids)
	} else {
		log.Println("Getting all products")
		res, err = s.service.GetProducts(ctx, r.Skip, r.Take)
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("no products found")
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []*pb.Products{}
	for _, p := range res {
		products = append(products, &pb.Products{
			ProductId:  p.ID,
			Title:      p.Name,
			Price:      p.Price,
			ImageURL:   p.ImageUrl,
			SellerName: p.SellerName,
		})
	}
	return &pb.GetProductsResponse{Products: products}, nil
}
