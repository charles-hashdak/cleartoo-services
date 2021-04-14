// order-service/main.go

package main

import(
	"context"
	"fmt"
	"sync"
	"log"
	"os"

	pb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	"github.com/micro/go-micro/v2"
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
	walletCollection := client.Database("cleartoo").Collection("wallet")
	transactionCollection := client.Database("cleartoo").Collection("transaction")

	repository := &MongoRepository{orderCollection, walletCollection, transactionCollection}

	cartClient := cartPb.NewCartService("cleartoo.cart", service.Client())

	mutex := sync.Mutex{}

	h := &handler{repository, cartClient, mutex}

	if err := pb.RegisterOrderServiceHandler(service.Server(), h); err != nil{
		fmt.Println(err)
	}

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}