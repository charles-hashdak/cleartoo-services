// order-service/main.go

package main

import(
	"context"

	pb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
)

type handler struct{
	repository
}

func (s *handler) Order(ctx context.Context, req *pb.OrderRequest, res *pb.OrderResponse) error {

	err := s.repository.CreateOrder(ctx, MarshalOrderRequest(req))

	if err != nil{
		return nil
	}

	res.Added = true

	return nil
}

func (s *handler) GetOrders(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	orders, err := s.repository.GetOrders(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Orders = UnmarshalOrderCollection(orders)
	return nil
}

func (s *handler) GetSales(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	orders, err := s.repository.GetOrders(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Orders = UnmarshalOrderCollection(orders)
	return nil
}

func (s *handler) GetSingleOrder(ctx context.Context, req *pb.GetSingleRequest, res *pb.GetSingleResponse) error {
	order, err := s.repository.GetOrder(ctx, MarshalGetSingleRequest(req))
	if err != nil {
		return err
	}
	res.Order = UnmarshalOrder(order)
	return nil
}