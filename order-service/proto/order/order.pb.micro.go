// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order/order.proto

package order

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

// Api Endpoints for OrderService service

func NewOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderService service

type OrderService interface {
	Order(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...client.CallOption) (*UpdateOrderStatusResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...client.CallOption) (*CancelOrderResponse, error)
	GetSales(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	GetOrders(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	GetSingleOrder(ctx context.Context, in *GetSingleRequest, opts ...client.CallOption) (*GetSingleResponse, error)
	GetWallet(ctx context.Context, in *GetWalletRequest, opts ...client.CallOption) (*GetWalletResponse, error)
	InitializeWallet(ctx context.Context, in *InitializeWalletRequest, opts ...client.CallOption) (*InitializeWalletResponse, error)
	UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...client.CallOption) (*UpdateWalletResponse, error)
	AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...client.CallOption) (*AddTransactionResponse, error)
	GetInTransitOrders(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Order(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Order", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...client.CallOption) (*UpdateOrderStatusResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.UpdateOrderStatus", in)
	out := new(UpdateOrderStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...client.CallOption) (*CancelOrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CancelOrder", in)
	out := new(CancelOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetSales(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetSales", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetOrders(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrders", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetSingleOrder(ctx context.Context, in *GetSingleRequest, opts ...client.CallOption) (*GetSingleResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetSingleOrder", in)
	out := new(GetSingleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetWallet(ctx context.Context, in *GetWalletRequest, opts ...client.CallOption) (*GetWalletResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetWallet", in)
	out := new(GetWalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) InitializeWallet(ctx context.Context, in *InitializeWalletRequest, opts ...client.CallOption) (*InitializeWalletResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.InitializeWallet", in)
	out := new(InitializeWalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...client.CallOption) (*UpdateWalletResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.UpdateWallet", in)
	out := new(UpdateWalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...client.CallOption) (*AddTransactionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AddTransaction", in)
	out := new(AddTransactionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetInTransitOrders(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetInTransitOrders", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	Order(context.Context, *OrderRequest, *OrderResponse) error
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest, *UpdateOrderStatusResponse) error
	CancelOrder(context.Context, *CancelOrderRequest, *CancelOrderResponse) error
	GetSales(context.Context, *GetRequest, *GetResponse) error
	GetOrders(context.Context, *GetRequest, *GetResponse) error
	GetSingleOrder(context.Context, *GetSingleRequest, *GetSingleResponse) error
	GetWallet(context.Context, *GetWalletRequest, *GetWalletResponse) error
	InitializeWallet(context.Context, *InitializeWalletRequest, *InitializeWalletResponse) error
	UpdateWallet(context.Context, *UpdateWalletRequest, *UpdateWalletResponse) error
	AddTransaction(context.Context, *AddTransactionRequest, *AddTransactionResponse) error
	GetInTransitOrders(context.Context, *GetRequest, *GetResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		Order(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, out *UpdateOrderStatusResponse) error
		CancelOrder(ctx context.Context, in *CancelOrderRequest, out *CancelOrderResponse) error
		GetSales(ctx context.Context, in *GetRequest, out *GetResponse) error
		GetOrders(ctx context.Context, in *GetRequest, out *GetResponse) error
		GetSingleOrder(ctx context.Context, in *GetSingleRequest, out *GetSingleResponse) error
		GetWallet(ctx context.Context, in *GetWalletRequest, out *GetWalletResponse) error
		InitializeWallet(ctx context.Context, in *InitializeWalletRequest, out *InitializeWalletResponse) error
		UpdateWallet(ctx context.Context, in *UpdateWalletRequest, out *UpdateWalletResponse) error
		AddTransaction(ctx context.Context, in *AddTransactionRequest, out *AddTransactionResponse) error
		GetInTransitOrders(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) Order(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.Order(ctx, in, out)
}

func (h *orderServiceHandler) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, out *UpdateOrderStatusResponse) error {
	return h.OrderServiceHandler.UpdateOrderStatus(ctx, in, out)
}

func (h *orderServiceHandler) CancelOrder(ctx context.Context, in *CancelOrderRequest, out *CancelOrderResponse) error {
	return h.OrderServiceHandler.CancelOrder(ctx, in, out)
}

func (h *orderServiceHandler) GetSales(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.OrderServiceHandler.GetSales(ctx, in, out)
}

func (h *orderServiceHandler) GetOrders(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.OrderServiceHandler.GetOrders(ctx, in, out)
}

func (h *orderServiceHandler) GetSingleOrder(ctx context.Context, in *GetSingleRequest, out *GetSingleResponse) error {
	return h.OrderServiceHandler.GetSingleOrder(ctx, in, out)
}

func (h *orderServiceHandler) GetWallet(ctx context.Context, in *GetWalletRequest, out *GetWalletResponse) error {
	return h.OrderServiceHandler.GetWallet(ctx, in, out)
}

func (h *orderServiceHandler) InitializeWallet(ctx context.Context, in *InitializeWalletRequest, out *InitializeWalletResponse) error {
	return h.OrderServiceHandler.InitializeWallet(ctx, in, out)
}

func (h *orderServiceHandler) UpdateWallet(ctx context.Context, in *UpdateWalletRequest, out *UpdateWalletResponse) error {
	return h.OrderServiceHandler.UpdateWallet(ctx, in, out)
}

func (h *orderServiceHandler) AddTransaction(ctx context.Context, in *AddTransactionRequest, out *AddTransactionResponse) error {
	return h.OrderServiceHandler.AddTransaction(ctx, in, out)
}

func (h *orderServiceHandler) GetInTransitOrders(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.OrderServiceHandler.GetInTransitOrders(ctx, in, out)
}
