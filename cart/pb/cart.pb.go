// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: cart.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string  `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Title     string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price     float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	SellerId  string  `protobuf:"bytes,4,opt,name=sellerId,proto3" json:"sellerId,omitempty"`
	Quantity  uint64  `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CartItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartItem) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

func (x *CartItem) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId     string      `protobuf:"bytes,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	Items      []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice float64     `protobuf:"fixed64,3,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *Cart) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Cart) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity  uint64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *AddToCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddToCartRequest) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFromCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type UpdateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId       string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	UpdatedQuantity uint64 `protobuf:"varint,2,opt,name=updatedQuantity,proto3" json:"updatedQuantity,omitempty"`
}

func (x *UpdateCartRequest) Reset() {
	*x = UpdateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartRequest) ProtoMessage() {}

func (x *UpdateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartRequest.ProtoReflect.Descriptor instead.
func (*UpdateCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateCartRequest) GetUpdatedQuantity() uint64 {
	if x != nil {
		return x.UpdatedQuantity
	}
	return 0
}

type IssueGuestCartTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestToken string `protobuf:"bytes,1,opt,name=guestToken,proto3" json:"guestToken,omitempty"`
}

func (x *IssueGuestCartTokenResponse) Reset() {
	*x = IssueGuestCartTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueGuestCartTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueGuestCartTokenResponse) ProtoMessage() {}

func (x *IssueGuestCartTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueGuestCartTokenResponse.ProtoReflect.Descriptor instead.
func (*IssueGuestCartTokenResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{5}
}

func (x *IssueGuestCartTokenResponse) GetGuestToken() string {
	if x != nil {
		return x.GuestToken
	}
	return ""
}

type ValidateGuestCartTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestId string `protobuf:"bytes,1,opt,name=guestId,proto3" json:"guestId,omitempty"`
}

func (x *ValidateGuestCartTokenResponse) Reset() {
	*x = ValidateGuestCartTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateGuestCartTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateGuestCartTokenResponse) ProtoMessage() {}

func (x *ValidateGuestCartTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateGuestCartTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateGuestCartTokenResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateGuestCartTokenResponse) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

type CartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{7}
}

func (x *CartResponse) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01,
	0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x62, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x4c, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x35,
	0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x3d, 0x0a, 0x1b, 0x49, 0x73, 0x73, 0x75, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x3a, 0x0a, 0x1e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x0c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x32, 0xe1, 0x03, 0x0a, 0x0b,
	0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x13, 0x49, 0x73, 0x73, 0x75, 0x65, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData = file_cart_proto_rawDesc
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_proto_rawDescData)
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cart_proto_goTypes = []any{
	(*CartItem)(nil),                       // 0: pb.CartItem
	(*Cart)(nil),                           // 1: pb.Cart
	(*AddToCartRequest)(nil),               // 2: pb.AddToCartRequest
	(*RemoveFromCartRequest)(nil),          // 3: pb.RemoveFromCartRequest
	(*UpdateCartRequest)(nil),              // 4: pb.UpdateCartRequest
	(*IssueGuestCartTokenResponse)(nil),    // 5: pb.IssueGuestCartTokenResponse
	(*ValidateGuestCartTokenResponse)(nil), // 6: pb.ValidateGuestCartTokenResponse
	(*CartResponse)(nil),                   // 7: pb.CartResponse
	(*emptypb.Empty)(nil),                  // 8: google.protobuf.Empty
}
var file_cart_proto_depIdxs = []int32{
	0, // 0: pb.Cart.items:type_name -> pb.CartItem
	1, // 1: pb.CartResponse.cart:type_name -> pb.Cart
	2, // 2: pb.CartService.AddItemToCart:input_type -> pb.AddToCartRequest
	3, // 3: pb.CartService.RemoveItemFromCart:input_type -> pb.RemoveFromCartRequest
	8, // 4: pb.CartService.GetCart:input_type -> google.protobuf.Empty
	4, // 5: pb.CartService.UpdateCart:input_type -> pb.UpdateCartRequest
	8, // 6: pb.CartService.DeleteCart:input_type -> google.protobuf.Empty
	8, // 7: pb.CartService.IssueGuestCartToken:input_type -> google.protobuf.Empty
	8, // 8: pb.CartService.ValidateGuestCartToken:input_type -> google.protobuf.Empty
	7, // 9: pb.CartService.AddItemToCart:output_type -> pb.CartResponse
	7, // 10: pb.CartService.RemoveItemFromCart:output_type -> pb.CartResponse
	7, // 11: pb.CartService.GetCart:output_type -> pb.CartResponse
	7, // 12: pb.CartService.UpdateCart:output_type -> pb.CartResponse
	7, // 13: pb.CartService.DeleteCart:output_type -> pb.CartResponse
	5, // 14: pb.CartService.IssueGuestCartToken:output_type -> pb.IssueGuestCartTokenResponse
	6, // 15: pb.CartService.ValidateGuestCartToken:output_type -> pb.ValidateGuestCartTokenResponse
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CartItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Cart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddToCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveFromCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*IssueGuestCartTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ValidateGuestCartTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_rawDesc = nil
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
