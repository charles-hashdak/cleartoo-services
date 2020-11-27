// cart-service/proto/cart/cart.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: proto/cart/cart.proto

package cart

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Products []*Product `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[0]
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
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{0}
}

func (x *Cart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cart) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Cart) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Disponible bool   `protobuf:"varint,2,opt,name=disponible,proto3" json:"disponible,omitempty"`
	Title      string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Price      int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Photo      *Photo `protobuf:"bytes,5,opt,name=photo,proto3" json:"photo,omitempty"`
	Category   string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Size       string `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
	Color1     *Color `protobuf:"bytes,8,opt,name=color1,proto3" json:"color1,omitempty"`
	Color2     *Color `protobuf:"bytes,9,opt,name=color2,proto3" json:"color2,omitempty"`
	Brand      string `protobuf:"bytes,10,opt,name=brand,proto3" json:"brand,omitempty"`
	Condition  string `protobuf:"bytes,11,opt,name=condition,proto3" json:"condition,omitempty"`
	Material   string `protobuf:"bytes,12,opt,name=material,proto3" json:"material,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetDisponible() bool {
	if x != nil {
		return x.Disponible
	}
	return false
}

func (x *Product) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Product) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetPhoto() *Photo {
	if x != nil {
		return x.Photo
	}
	return nil
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *Product) GetColor1() *Color {
	if x != nil {
		return x.Color1
	}
	return nil
}

func (x *Product) GetColor2() *Color {
	if x != nil {
		return x.Color2
	}
	return nil
}

func (x *Product) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Product) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *Product) GetMaterial() string {
	if x != nil {
		return x.Material
	}
	return ""
}

type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	IsMain bool   `protobuf:"varint,3,opt,name=is_main,json=isMain,proto3" json:"is_main,omitempty"`
	Height int32  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Width  int32  `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *Photo) Reset() {
	*x = Photo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{2}
}

func (x *Photo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Photo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Photo) GetIsMain() bool {
	if x != nil {
		return x.IsMain
	}
	return false
}

func (x *Photo) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Photo) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	HexCode string `protobuf:"bytes,3,opt,name=hex_code,json=hexCode,proto3" json:"hex_code,omitempty"`
	Image   string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{3}
}

func (x *Color) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Color) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Color) GetHexCode() string {
	if x != nil {
		return x.HexCode
	}
	return ""
}

func (x *Color) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Product *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[4]
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
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{4}
}

func (x *AddToCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddToCartRequest) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type AddToCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Added bool `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{5}
}

func (x *AddToCartResponse) GetAdded() bool {
	if x != nil {
		return x.Added
	}
	return false
}

type DeleteFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *DeleteFromCartRequest) Reset() {
	*x = DeleteFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFromCartRequest) ProtoMessage() {}

func (x *DeleteFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[6]
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
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFromCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteFromCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type DeleteFromCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteFromCartResponse) Reset() {
	*x = DeleteFromCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFromCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFromCartResponse) ProtoMessage() {}

func (x *DeleteFromCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFromCartResponse.ProtoReflect.Descriptor instead.
func (*DeleteFromCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFromCartResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{8}
}

func (x *GetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{9}
}

func (x *GetResponse) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type IsInCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *IsInCartRequest) Reset() {
	*x = IsInCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsInCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsInCartRequest) ProtoMessage() {}

func (x *IsInCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsInCartRequest.ProtoReflect.Descriptor instead.
func (*IsInCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{10}
}

func (x *IsInCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IsInCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type IsInCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In bool `protobuf:"varint,1,opt,name=in,proto3" json:"in,omitempty"`
}

func (x *IsInCartResponse) Reset() {
	*x = IsInCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_cart_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsInCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsInCartResponse) ProtoMessage() {}

func (x *IsInCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_cart_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsInCartResponse.ProtoReflect.Descriptor instead.
func (*IsInCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_cart_proto_rawDescGZIP(), []int{11}
}

func (x *IsInCartResponse) GetIn() bool {
	if x != nil {
		return x.In
	}
	return false
}

var File_proto_cart_cart_proto protoreflect.FileDescriptor

var file_proto_cart_cart_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x5a, 0x0a,
	0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xd2, 0x02, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69,
	0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x6f,
	0x6e, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x31, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x31, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x32, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x70,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4d, 0x61,
	0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x22, 0x5c, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x68, 0x65, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x68, 0x65, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x54,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x22,
	0x4f, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x22, 0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x49, 0x0a, 0x0f, 0x49, 0x73,
	0x49, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x49, 0x73, 0x49, 0x6e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x69, 0x6e, 0x32, 0x85, 0x03, 0x0a, 0x0b, 0x43, 0x61,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72,
	0x74, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x49, 0x73, 0x49, 0x6e, 0x43, 0x61, 0x72,
	0x74, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x49, 0x73, 0x49, 0x6e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x49, 0x73, 0x49, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cart_cart_proto_rawDescOnce sync.Once
	file_proto_cart_cart_proto_rawDescData = file_proto_cart_cart_proto_rawDesc
)

func file_proto_cart_cart_proto_rawDescGZIP() []byte {
	file_proto_cart_cart_proto_rawDescOnce.Do(func() {
		file_proto_cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cart_cart_proto_rawDescData)
	})
	return file_proto_cart_cart_proto_rawDescData
}

var file_proto_cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_cart_cart_proto_goTypes = []interface{}{
	(*Cart)(nil),                   // 0: cart.Cart
	(*Product)(nil),                // 1: cart.Product
	(*Photo)(nil),                  // 2: cart.Photo
	(*Color)(nil),                  // 3: cart.Color
	(*AddToCartRequest)(nil),       // 4: cart.AddToCartRequest
	(*AddToCartResponse)(nil),      // 5: cart.AddToCartResponse
	(*DeleteFromCartRequest)(nil),  // 6: cart.DeleteFromCartRequest
	(*DeleteFromCartResponse)(nil), // 7: cart.DeleteFromCartResponse
	(*GetRequest)(nil),             // 8: cart.GetRequest
	(*GetResponse)(nil),            // 9: cart.GetResponse
	(*IsInCartRequest)(nil),        // 10: cart.IsInCartRequest
	(*IsInCartResponse)(nil),       // 11: cart.IsInCartResponse
}
var file_proto_cart_cart_proto_depIdxs = []int32{
	1,  // 0: cart.Cart.products:type_name -> cart.Product
	2,  // 1: cart.Product.photo:type_name -> cart.Photo
	3,  // 2: cart.Product.color1:type_name -> cart.Color
	3,  // 3: cart.Product.color2:type_name -> cart.Color
	1,  // 4: cart.AddToCartRequest.product:type_name -> cart.Product
	0,  // 5: cart.GetResponse.cart:type_name -> cart.Cart
	8,  // 6: cart.CartService.CreateCart:input_type -> cart.GetRequest
	4,  // 7: cart.CartService.AddToCart:input_type -> cart.AddToCartRequest
	6,  // 8: cart.CartService.DeleteFromCart:input_type -> cart.DeleteFromCartRequest
	8,  // 9: cart.CartService.GetCart:input_type -> cart.GetRequest
	10, // 10: cart.CartService.IsInCart:input_type -> cart.IsInCartRequest
	8,  // 11: cart.CartService.EmptyCart:input_type -> cart.GetRequest
	5,  // 12: cart.CartService.CreateCart:output_type -> cart.AddToCartResponse
	5,  // 13: cart.CartService.AddToCart:output_type -> cart.AddToCartResponse
	7,  // 14: cart.CartService.DeleteFromCart:output_type -> cart.DeleteFromCartResponse
	9,  // 15: cart.CartService.GetCart:output_type -> cart.GetResponse
	11, // 16: cart.CartService.IsInCart:output_type -> cart.IsInCartResponse
	7,  // 17: cart.CartService.EmptyCart:output_type -> cart.DeleteFromCartResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_cart_cart_proto_init() }
func file_proto_cart_cart_proto_init() {
	if File_proto_cart_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cart_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_cart_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_cart_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Photo); i {
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
		file_proto_cart_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_proto_cart_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_cart_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponse); i {
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
		file_proto_cart_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_cart_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFromCartResponse); i {
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
		file_proto_cart_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_cart_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_proto_cart_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsInCartRequest); i {
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
		file_proto_cart_cart_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsInCartResponse); i {
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
			RawDescriptor: file_proto_cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cart_cart_proto_goTypes,
		DependencyIndexes: file_proto_cart_cart_proto_depIdxs,
		MessageInfos:      file_proto_cart_cart_proto_msgTypes,
	}.Build()
	File_proto_cart_cart_proto = out.File
	file_proto_cart_cart_proto_rawDesc = nil
	file_proto_cart_cart_proto_goTypes = nil
	file_proto_cart_cart_proto_depIdxs = nil
}
