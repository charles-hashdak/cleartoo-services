// cart-service/main.go

package main

import(
	"context"

	pb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
)

type handler struct{
	repository
}

func (s *handler) AddToCart(ctx context.Context, req *pb.AddToCartRequest, res *pb.AddToCartResponse) error {

	_, err := s.repository.AddToCart(ctx, req)

	if err != nil{
		return nil
	}

	res.Added = true

	return nil
}

func (s *handler) DeleteFromCart(ctx context.Context, req *pb.DeleteFromCartRequest, res *pb.DeleteFromCartResponse) error {

	_, err := s.repository.DeleteFromCart(ctx, req)

	if err != nil{
		return nil
	}

	res.Deleted = true

	return nil
}

func (s *handler) GetCart(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	cart, err := s.repository.GetCart(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res = UnmarshalCart(cart)
	return nil
}