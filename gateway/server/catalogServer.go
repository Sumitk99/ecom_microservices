package server

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"log"
)

func (s *Server) GetProduct(ctx context.Context, productId string) (*models.Product, error) {
	log.Println("Getting product")
	res, err := s.CatalogClient.GetProduct(ctx, &pb.GetProductRequest{Id: productId})
	log.Println("Got")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	clrs := []models.Color{}
	for _, clr := range res.Product.Colors {
		clrs = append(clrs, models.Color{
			ColorName: clr.ColorName,
			Hex:       clr.Hex,
		})
	}
	return &models.Product{
		ID:          res.Product.ProductId,
		Name:        res.Product.Title,
		Description: res.Product.Description,
		Price:       res.Product.Price,
		Stock:       res.Product.AvailableQuantity,
		ImageUrl:    res.Product.ImageURL,
		Category:    res.Product.Category,
		SellerID:    res.Product.SellerId,
		SellerName:  res.Product.SellerName,
		Locations:   res.Product.Locations,
		Sizes:       res.Product.Sizes,
		Colors:      clrs,
	}, nil
}

func (s *Server) GetProducts(ctx context.Context, search string, skip, take int) (*models.GetProductsResponse, error) {
	res, err := s.CatalogClient.GetProducts(
		ctx,
		&pb.GetProductsRequest{
			Query: search,
			Skip:  uint64(skip),
			Take:  uint64(take),
		},
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []models.Products{}
	for _, product := range res.Products {
		products = append(products, models.Products{
			ProductId:  product.ProductId,
			Title:      product.Title,
			Price:      product.Price,
			SellerName: product.SellerName,
			ImageURL:   product.ImageURL,
		})
	}
	return &models.GetProductsResponse{Products: products}, nil
}

func (s *Server) PostProduct(ctx context.Context, product *models.ProductDocument) (*pb.PostProductResponse, error) {
	clrs := []*pb.Color{}
	for _, clr := range product.Colors {
		clrs = append(clrs, &pb.Color{
			ColorName: clr.ColorName,
			Hex:       clr.Hex,
		})
	}
	res, err := s.CatalogClient.PostProduct(ctx, &pb.PostProductRequest{
		Title:             product.Name,
		Description:       product.Description,
		Price:             product.Price,
		SellerId:          product.SellerID,
		SellerName:        product.SellerName,
		ImageURL:          product.ImageUrl,
		Category:          product.Category,
		AvailableQuantity: product.Stock,
		Locations:         product.Locations,
		Sizes:             product.Sizes,
		Colors:            clrs,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
	//clrs = []models.Color{}
	//for _, color := range res.Product.Colors {
	//	clrs = append(clrs, models.Color{
	//		ColorName: color.ColorName,
	//		Hex:       color.Hex,
	//	})
	//}
	//return &models.Product{
	//	ID:          res.Product.ProductId,
	//	Name:        res.Product.Title,
	//	Description: res.Product.Description,
	//	Price:       res.Product.Price,
	//	Category:    res.Product.Category,
	//	Stock:       res.Product.AvailableQuantity,
	//	Locations:   res.Product.Locations,
	//	ImageUrl:    res.Product.ImageURL,
	//	Sizes:       res.Product.Sizes,
	//	Colors:      clrs,
	//	SellerID:    res.Product.SellerId,
	//	SellerName:  res.Product.SellerName,
	//}, nil
}
