// cart-service/main.go

package main

import(
	"context"
	_ "fmt"

	pb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
)

type handler struct{
	repository
}

func (s *handler) CreateCart(ctx context.Context, req *pb.GetRequest, res *pb.AddToCartResponse) error {

	err := s.repository.CreateCart(ctx, MarshalGetRequest(req))

	if err != nil{
		return nil
	}

	res.Added = true

	return nil
}

func (s *handler) AddToCart(ctx context.Context, req *pb.AddToCartRequest, res *pb.AddToCartResponse) error {
	err := s.repository.AddToCart(ctx, MarshalAddToCartRequest(req))

	if err != nil{
		return nil
	}

	res.Added = true

	return nil
}

func (s *handler) DeleteFromCart(ctx context.Context, req *pb.DeleteFromCartRequest, res *pb.DeleteFromCartResponse) error {

	err := s.repository.DeleteFromCart(ctx, MarshalDeleteFromCartRequest(req))

	if err != nil{
		return nil
	}

	res.Deleted = true

	return nil
}

func (s *handler) EmptyCart(ctx context.Context, req *pb.GetRequest, res *pb.DeleteFromCartResponse) error {

	err := s.repository.EmptyCart(ctx, MarshalGetRequest(req))

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
	res.Cart = UnmarshalCart(cart)
	return nil
}

func (s *handler) IsInCart(ctx context.Context, req *pb.IsInCartRequest, res *pb.IsInCartResponse) error {
	isincart, err := s.repository.IsInCart(ctx, MarshalIsInCartRequest(req))

	if err != nil{
		return nil
	}

	res.In = isincart

	return nil
}