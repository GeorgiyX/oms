// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: checkout.proto

package checkout

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64  `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku   uint32 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[0]
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
	return file_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *AddToCartRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *AddToCartRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64  `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku   uint32 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *DeleteFromCartRequest) Reset() {
	*x = DeleteFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFromCartRequest) ProtoMessage() {}

func (x *DeleteFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFromCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteFromCartRequest) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFromCartRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *DeleteFromCartRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *DeleteFromCartRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListCartRequest) Reset() {
	*x = ListCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartRequest) ProtoMessage() {}

func (x *ListCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartRequest.ProtoReflect.Descriptor instead.
func (*ListCartRequest) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{2}
}

func (x *ListCartRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

type ListCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*ListCartResponse_CartItemResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice uint32                               `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *ListCartResponse) Reset() {
	*x = ListCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartResponse) ProtoMessage() {}

func (x *ListCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartResponse.ProtoReflect.Descriptor instead.
func (*ListCartResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{3}
}

func (x *ListCartResponse) GetItems() []*ListCartResponse_CartItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListCartResponse) GetTotalPrice() uint32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type PurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{4}
}

func (x *PurchaseRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

type PurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{5}
}

func (x *PurchaseResponse) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type ListCartResponse_CartItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ListCartResponse_CartItemResponse) Reset() {
	*x = ListCartResponse_CartItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartResponse_CartItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartResponse_CartItemResponse) ProtoMessage() {}

func (x *ListCartResponse_CartItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartResponse_CartItemResponse.ProtoReflect.Descriptor instead.
func (*ListCartResponse_CartItemResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListCartResponse_CartItemResponse) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *ListCartResponse_CartItemResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListCartResponse_CartItemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListCartResponse_CartItemResponse) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_checkout_proto protoreflect.FileDescriptor

var file_checkout_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b,
	0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0xdc, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x64, 0x0a, 0x10,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x10, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0x9c, 0x02, 0x0a, 0x08, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x41, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x19, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x19, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x32, 0x35, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x3b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkout_proto_rawDescOnce sync.Once
	file_checkout_proto_rawDescData = file_checkout_proto_rawDesc
)

func file_checkout_proto_rawDescGZIP() []byte {
	file_checkout_proto_rawDescOnce.Do(func() {
		file_checkout_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkout_proto_rawDescData)
	})
	return file_checkout_proto_rawDescData
}

var file_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_checkout_proto_goTypes = []interface{}{
	(*AddToCartRequest)(nil),                  // 0: checkout.AddToCartRequest
	(*DeleteFromCartRequest)(nil),             // 1: checkout.DeleteFromCartRequest
	(*ListCartRequest)(nil),                   // 2: checkout.ListCartRequest
	(*ListCartResponse)(nil),                  // 3: checkout.ListCartResponse
	(*PurchaseRequest)(nil),                   // 4: checkout.PurchaseRequest
	(*PurchaseResponse)(nil),                  // 5: checkout.PurchaseResponse
	(*ListCartResponse_CartItemResponse)(nil), // 6: checkout.ListCartResponse.CartItemResponse
	(*emptypb.Empty)(nil),                     // 7: google.protobuf.Empty
}
var file_checkout_proto_depIdxs = []int32{
	6, // 0: checkout.ListCartResponse.items:type_name -> checkout.ListCartResponse.CartItemResponse
	0, // 1: checkout.Checkout.AddToCart:input_type -> checkout.AddToCartRequest
	1, // 2: checkout.Checkout.DeleteFromCart:input_type -> checkout.DeleteFromCartRequest
	2, // 3: checkout.Checkout.ListCart:input_type -> checkout.ListCartRequest
	4, // 4: checkout.Checkout.Purchase:input_type -> checkout.PurchaseRequest
	7, // 5: checkout.Checkout.AddToCart:output_type -> google.protobuf.Empty
	7, // 6: checkout.Checkout.DeleteFromCart:output_type -> google.protobuf.Empty
	3, // 7: checkout.Checkout.ListCart:output_type -> checkout.ListCartResponse
	5, // 8: checkout.Checkout.Purchase:output_type -> checkout.PurchaseResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_checkout_proto_init() }
func file_checkout_proto_init() {
	if File_checkout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_checkout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFromCartRequest); i {
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
		file_checkout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartRequest); i {
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
		file_checkout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartResponse); i {
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
		file_checkout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseRequest); i {
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
		file_checkout_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseResponse); i {
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
		file_checkout_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartResponse_CartItemResponse); i {
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
			RawDescriptor: file_checkout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkout_proto_goTypes,
		DependencyIndexes: file_checkout_proto_depIdxs,
		MessageInfos:      file_checkout_proto_msgTypes,
	}.Build()
	File_checkout_proto = out.File
	file_checkout_proto_rawDesc = nil
	file_checkout_proto_goTypes = nil
	file_checkout_proto_depIdxs = nil
}
