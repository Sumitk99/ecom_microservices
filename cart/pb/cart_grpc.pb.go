// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: cart.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CartService_AddItemToCart_FullMethodName          = "/pb.CartService/AddItemToCart"
	CartService_RemoveItemFromCart_FullMethodName     = "/pb.CartService/RemoveItemFromCart"
	CartService_GetCart_FullMethodName                = "/pb.CartService/GetCart"
	CartService_UpdateCart_FullMethodName             = "/pb.CartService/UpdateCart"
	CartService_DeleteCart_FullMethodName             = "/pb.CartService/DeleteCart"
	CartService_IssueGuestCartToken_FullMethodName    = "/pb.CartService/IssueGuestCartToken"
	CartService_ValidateGuestCartToken_FullMethodName = "/pb.CartService/ValidateGuestCartToken"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	AddItemToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	RemoveItemFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	GetCart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CartResponse, error)
	UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	DeleteCart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CartResponse, error)
	IssueGuestCartToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IssueGuestCartTokenResponse, error)
	ValidateGuestCartToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ValidateGuestCartTokenResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddItemToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_AddItemToCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveItemFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_RemoveItemFromCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_UpdateCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteCart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, CartService_DeleteCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) IssueGuestCartToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IssueGuestCartTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IssueGuestCartTokenResponse)
	err := c.cc.Invoke(ctx, CartService_IssueGuestCartToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ValidateGuestCartToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ValidateGuestCartTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateGuestCartTokenResponse)
	err := c.cc.Invoke(ctx, CartService_ValidateGuestCartToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility.
type CartServiceServer interface {
	AddItemToCart(context.Context, *AddToCartRequest) (*CartResponse, error)
	RemoveItemFromCart(context.Context, *RemoveFromCartRequest) (*CartResponse, error)
	GetCart(context.Context, *emptypb.Empty) (*CartResponse, error)
	UpdateCart(context.Context, *UpdateCartRequest) (*CartResponse, error)
	DeleteCart(context.Context, *emptypb.Empty) (*CartResponse, error)
	IssueGuestCartToken(context.Context, *emptypb.Empty) (*IssueGuestCartTokenResponse, error)
	ValidateGuestCartToken(context.Context, *emptypb.Empty) (*ValidateGuestCartTokenResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCartServiceServer struct{}

func (UnimplementedCartServiceServer) AddItemToCart(context.Context, *AddToCartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItemToCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveItemFromCart(context.Context, *RemoveFromCartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItemFromCart not implemented")
}
func (UnimplementedCartServiceServer) GetCart(context.Context, *emptypb.Empty) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) UpdateCart(context.Context, *UpdateCartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedCartServiceServer) DeleteCart(context.Context, *emptypb.Empty) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedCartServiceServer) IssueGuestCartToken(context.Context, *emptypb.Empty) (*IssueGuestCartTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueGuestCartToken not implemented")
}
func (UnimplementedCartServiceServer) ValidateGuestCartToken(context.Context, *emptypb.Empty) (*ValidateGuestCartTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateGuestCartToken not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}
func (UnimplementedCartServiceServer) testEmbeddedByValue()                     {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	// If the following call pancis, it indicates UnimplementedCartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_AddItemToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddItemToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_AddItemToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddItemToCart(ctx, req.(*AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveItemFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveItemFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_RemoveItemFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveItemFromCart(ctx, req.(*RemoveFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_UpdateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).UpdateCart(ctx, req.(*UpdateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_DeleteCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DeleteCart(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_IssueGuestCartToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).IssueGuestCartToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_IssueGuestCartToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).IssueGuestCartToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ValidateGuestCartToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ValidateGuestCartToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_ValidateGuestCartToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ValidateGuestCartToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItemToCart",
			Handler:    _CartService_AddItemToCart_Handler,
		},
		{
			MethodName: "RemoveItemFromCart",
			Handler:    _CartService_RemoveItemFromCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _CartService_UpdateCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _CartService_DeleteCart_Handler,
		},
		{
			MethodName: "IssueGuestCartToken",
			Handler:    _CartService_IssueGuestCartToken_Handler,
		},
		{
			MethodName: "ValidateGuestCartToken",
			Handler:    _CartService_ValidateGuestCartToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
