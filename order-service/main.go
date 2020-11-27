// order-service/main.go

package main

import(
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	"github.com/micro/go-micro/v2"
	_ "github.com/asim/nitro-plugins/registry/mdns"
)

func main(){

	service := micro.NewService(
		micro.Name("cleartoo.order"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	orderCollection := client.Database("cleartoo").Collection("order")

	repository := &MongoRepository{orderCollection}

	cartClient := cartPb.NewCartService("cleartoo.cart", service.Client())

	h := &handler{repository, cartClient}

	if err := pb.RegisterOrderServiceHandler(service.Server(), h); err != nil{
		fmt.Println(err)
	}

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}