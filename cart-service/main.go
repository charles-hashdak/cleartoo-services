// cart-service/main.go

package main

import(
	"context"
	"log"

	pb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	"github.com/micro/go-micro/v2"
)

type repository interface{
	AddToCart(*pb.AddToCartRequest) (*pb.AddToCartRequest, error)
}

type Repository struct{
	mu			sync.RWMutex
	items		[]*pb.AddToCartRequest
}

func (repo *Repository) AddToCart(req *pb.AddToCartRequest) (*pb.AddToCartRequest, error){
	repo.mu.Lock()
	updated := append(repo.items, req)
	repo.items = updated
	repo.mu.Unlock()
	return req, nil
}

type service struct{
	repo repository
}

func (s *service) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error){

	item, err := s.repo.AddToCart(req)

	if err != nil{
		return nil, err
	}

	return &pb.AddToCartResponse{Added: true}, nil
}

func main(){

	repo := &Repository{}

	service := micro.NewService(
		micro.Name("cleartoo.service.cart"),
	)

	service.Init()

	if err := pb.RegisterCartServiceHandler(service.Server(), &cartService{repo}); err != nil{
		log.Panic(err)
	}

	if err := service.Run(); err != nil{
		log.Panic(err)
	}
}