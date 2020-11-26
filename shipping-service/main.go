// shipping-service/main.go

package main

import(
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/charles-hashdak/cleartoo-services/shipping-service/proto/shipping"
	"github.com/micro/go-micro/v2"
	_ "github.com/asim/nitro-plugins/registry/mdns"
)

func main(){

	service := micro.NewService(
		micro.Name("cleartoo.shipping"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	addressesCollection := client.Database("cleartoo").Collection("addresses")
	shipmentsCollection := client.Database("cleartoo").Collection("shipments")
	methodsCollection := client.Database("cleartoo").Collection("methods")
	countriesCollection := client.Database("cleartoo").Collection("countries")
	citiesCollection := client.Database("cleartoo").Collection("cities")

	repository := &MongoRepository{addressesCollection, shipmentsCollection, methodsCollection, countriesCollection, citiesCollection}

	h := &handler{repository}

	if err := pb.RegistershippingServiceHandler(service.Server(), h); err != nil{
		fmt.Println(err)
	}

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}