// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/catalog/catalog.proto

package catalog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CatalogService service

func NewCatalogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CatalogService service

type CatalogService interface {
	CreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*CreateProductResponse, error)
	EditProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*EditProductResponse, error)
	Unavailable(ctx context.Context, in *Product, opts ...client.CallOption) (*EditProductResponse, error)
	Available(ctx context.Context, in *Product, opts ...client.CallOption) (*EditProductResponse, error)
	CreateOffer(ctx context.Context, in *CreateOfferRequest, opts ...client.CallOption) (*CreateOfferResponse, error)
	EditOffer(ctx context.Context, in *Offer, opts ...client.CallOption) (*EditOfferResponse, error)
	GetProducts(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetProductsResponse, error)
	GetProduct(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Product, error)
	Wish(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*WishResponse, error)
	GetWishes(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetProductsResponse, error)
	GetSizes(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetSizesResponse, error)
	GetGenders(ctx context.Context, in *Request, opts ...client.CallOption) (*GetGendersResponse, error)
	GetCategories(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetCategoriesResponse, error)
	GetBrands(ctx context.Context, in *Request, opts ...client.CallOption) (*GetBrandsResponse, error)
	GetColors(ctx context.Context, in *Request, opts ...client.CallOption) (*GetColorsResponse, error)
	GetConditions(ctx context.Context, in *Request, opts ...client.CallOption) (*GetConditionsResponse, error)
	GetAddProductData(ctx context.Context, in *Request, opts ...client.CallOption) (*GetAddProductDataResponse, error)
}

type catalogService struct {
	c    client.Client
	name string
}

func NewCatalogService(name string, c client.Client) CatalogService {
	return &catalogService{
		c:    c,
		name: name,
	}
}

func (c *catalogService) CreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*CreateProductResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.CreateProduct", in)
	out := new(CreateProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) EditProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*EditProductResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.EditProduct", in)
	out := new(EditProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) Unavailable(ctx context.Context, in *Product, opts ...client.CallOption) (*EditProductResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.Unavailable", in)
	out := new(EditProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) Available(ctx context.Context, in *Product, opts ...client.CallOption) (*EditProductResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.Available", in)
	out := new(EditProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) CreateOffer(ctx context.Context, in *CreateOfferRequest, opts ...client.CallOption) (*CreateOfferResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.CreateOffer", in)
	out := new(CreateOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) EditOffer(ctx context.Context, in *Offer, opts ...client.CallOption) (*EditOfferResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.EditOffer", in)
	out := new(EditOfferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetProducts(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetProductsResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetProducts", in)
	out := new(GetProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetProduct(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetProduct", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) Wish(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*WishResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.Wish", in)
	out := new(WishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetWishes(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetProductsResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetWishes", in)
	out := new(GetProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetSizes(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetSizesResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetSizes", in)
	out := new(GetSizesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetGenders(ctx context.Context, in *Request, opts ...client.CallOption) (*GetGendersResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetGenders", in)
	out := new(GetGendersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetCategories(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetCategoriesResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetCategories", in)
	out := new(GetCategoriesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetBrands(ctx context.Context, in *Request, opts ...client.CallOption) (*GetBrandsResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetBrands", in)
	out := new(GetBrandsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetColors(ctx context.Context, in *Request, opts ...client.CallOption) (*GetColorsResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetColors", in)
	out := new(GetColorsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetConditions(ctx context.Context, in *Request, opts ...client.CallOption) (*GetConditionsResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetConditions", in)
	out := new(GetConditionsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetAddProductData(ctx context.Context, in *Request, opts ...client.CallOption) (*GetAddProductDataResponse, error) {
	req := c.c.NewRequest(c.name, "CatalogService.GetAddProductData", in)
	out := new(GetAddProductDataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CatalogService service

type CatalogServiceHandler interface {
	CreateProduct(context.Context, *Product, *CreateProductResponse) error
	EditProduct(context.Context, *Product, *EditProductResponse) error
	Unavailable(context.Context, *Product, *EditProductResponse) error
	Available(context.Context, *Product, *EditProductResponse) error
	CreateOffer(context.Context, *CreateOfferRequest, *CreateOfferResponse) error
	EditOffer(context.Context, *Offer, *EditOfferResponse) error
	GetProducts(context.Context, *GetRequest, *GetProductsResponse) error
	GetProduct(context.Context, *GetRequest, *Product) error
	Wish(context.Context, *GetRequest, *WishResponse) error
	GetWishes(context.Context, *GetRequest, *GetProductsResponse) error
	GetSizes(context.Context, *GetRequest, *GetSizesResponse) error
	GetGenders(context.Context, *Request, *GetGendersResponse) error
	GetCategories(context.Context, *GetRequest, *GetCategoriesResponse) error
	GetBrands(context.Context, *Request, *GetBrandsResponse) error
	GetColors(context.Context, *Request, *GetColorsResponse) error
	GetConditions(context.Context, *Request, *GetConditionsResponse) error
	GetAddProductData(context.Context, *Request, *GetAddProductDataResponse) error
}

func RegisterCatalogServiceHandler(s server.Server, hdlr CatalogServiceHandler, opts ...server.HandlerOption) error {
	type catalogService interface {
		CreateProduct(ctx context.Context, in *Product, out *CreateProductResponse) error
		EditProduct(ctx context.Context, in *Product, out *EditProductResponse) error
		Unavailable(ctx context.Context, in *Product, out *EditProductResponse) error
		Available(ctx context.Context, in *Product, out *EditProductResponse) error
		CreateOffer(ctx context.Context, in *CreateOfferRequest, out *CreateOfferResponse) error
		EditOffer(ctx context.Context, in *Offer, out *EditOfferResponse) error
		GetProducts(ctx context.Context, in *GetRequest, out *GetProductsResponse) error
		GetProduct(ctx context.Context, in *GetRequest, out *Product) error
		Wish(ctx context.Context, in *GetRequest, out *WishResponse) error
		GetWishes(ctx context.Context, in *GetRequest, out *GetProductsResponse) error
		GetSizes(ctx context.Context, in *GetRequest, out *GetSizesResponse) error
		GetGenders(ctx context.Context, in *Request, out *GetGendersResponse) error
		GetCategories(ctx context.Context, in *GetRequest, out *GetCategoriesResponse) error
		GetBrands(ctx context.Context, in *Request, out *GetBrandsResponse) error
		GetColors(ctx context.Context, in *Request, out *GetColorsResponse) error
		GetConditions(ctx context.Context, in *Request, out *GetConditionsResponse) error
		GetAddProductData(ctx context.Context, in *Request, out *GetAddProductDataResponse) error
	}
	type CatalogService struct {
		catalogService
	}
	h := &catalogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CatalogService{h}, opts...))
}

type catalogServiceHandler struct {
	CatalogServiceHandler
}

func (h *catalogServiceHandler) CreateProduct(ctx context.Context, in *Product, out *CreateProductResponse) error {
	return h.CatalogServiceHandler.CreateProduct(ctx, in, out)
}

func (h *catalogServiceHandler) EditProduct(ctx context.Context, in *Product, out *EditProductResponse) error {
	return h.CatalogServiceHandler.EditProduct(ctx, in, out)
}

func (h *catalogServiceHandler) Unavailable(ctx context.Context, in *Product, out *EditProductResponse) error {
	return h.CatalogServiceHandler.Unavailable(ctx, in, out)
}

func (h *catalogServiceHandler) Available(ctx context.Context, in *Product, out *EditProductResponse) error {
	return h.CatalogServiceHandler.Available(ctx, in, out)
}

func (h *catalogServiceHandler) CreateOffer(ctx context.Context, in *CreateOfferRequest, out *CreateOfferResponse) error {
	return h.CatalogServiceHandler.CreateOffer(ctx, in, out)
}

func (h *catalogServiceHandler) EditOffer(ctx context.Context, in *Offer, out *EditOfferResponse) error {
	return h.CatalogServiceHandler.EditOffer(ctx, in, out)
}

func (h *catalogServiceHandler) GetProducts(ctx context.Context, in *GetRequest, out *GetProductsResponse) error {
	return h.CatalogServiceHandler.GetProducts(ctx, in, out)
}

func (h *catalogServiceHandler) GetProduct(ctx context.Context, in *GetRequest, out *Product) error {
	return h.CatalogServiceHandler.GetProduct(ctx, in, out)
}

func (h *catalogServiceHandler) Wish(ctx context.Context, in *GetRequest, out *WishResponse) error {
	return h.CatalogServiceHandler.Wish(ctx, in, out)
}

func (h *catalogServiceHandler) GetWishes(ctx context.Context, in *GetRequest, out *GetProductsResponse) error {
	return h.CatalogServiceHandler.GetWishes(ctx, in, out)
}

func (h *catalogServiceHandler) GetSizes(ctx context.Context, in *GetRequest, out *GetSizesResponse) error {
	return h.CatalogServiceHandler.GetSizes(ctx, in, out)
}

func (h *catalogServiceHandler) GetGenders(ctx context.Context, in *Request, out *GetGendersResponse) error {
	return h.CatalogServiceHandler.GetGenders(ctx, in, out)
}

func (h *catalogServiceHandler) GetCategories(ctx context.Context, in *GetRequest, out *GetCategoriesResponse) error {
	return h.CatalogServiceHandler.GetCategories(ctx, in, out)
}

func (h *catalogServiceHandler) GetBrands(ctx context.Context, in *Request, out *GetBrandsResponse) error {
	return h.CatalogServiceHandler.GetBrands(ctx, in, out)
}

func (h *catalogServiceHandler) GetColors(ctx context.Context, in *Request, out *GetColorsResponse) error {
	return h.CatalogServiceHandler.GetColors(ctx, in, out)
}

func (h *catalogServiceHandler) GetConditions(ctx context.Context, in *Request, out *GetConditionsResponse) error {
	return h.CatalogServiceHandler.GetConditions(ctx, in, out)
}

func (h *catalogServiceHandler) GetAddProductData(ctx context.Context, in *Request, out *GetAddProductDataResponse) error {
	return h.CatalogServiceHandler.GetAddProductData(ctx, in, out)
}
