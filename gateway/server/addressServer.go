package server

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetAddresses(ctx context.Context) (*pb.Addresses, error) {
	addresses, err := s.AccountClient.GetAddresses(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (s *Server) AddAddress(ctx context.Context, address *models.AddAddressRequest) (*pb.Address, error) {
	res, err := s.AccountClient.AddAddress(ctx, &pb.AddAddressRequest{
		Name:          address.Name,
		Phone:         address.Phone,
		ZipCode:       address.ZipCode,
		ApartmentUnit: address.ApartmentUnit,
		City:          address.City,
		State:         address.State,
		Country:       address.Country,
		IsDefault:     address.IsDefault,
		Street:        address.Street,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server) DeleteAddress(ctx context.Context, addressId string) error {
	_, err := s.AccountClient.DeleteAddress(ctx, &pb.DeleteAddressRequest{
		AddressId: addressId,
	})
	return err
}

func (s *Server) GetAddress(ctx context.Context, addressId string) (*pb.Address, error) {
	address, err := s.AccountClient.GetAddress(ctx, &pb.GetAddressRequest{
		AddressId: addressId,
	})
	if err != nil {
		return nil, err
	}
	return address, nil
}
