// catalog-service/main.go

package main

import(
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	"github.com/micro/go-micro/v2"
)

func main(){

	service := micro.NewService(
		micro.Name("cleartoo.catalog"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	productCollection := client.Database("cleartoo").Collection("products")
	sizeCollection := client.Database("cleartoo").Collection("sizes")
	genderCollection := client.Database("cleartoo").Collection("genders")
	categoryCollection := client.Database("cleartoo").Collection("categories")
	brandCollection := client.Database("cleartoo").Collection("brands")
	colorCollection := client.Database("cleartoo").Collection("colors")
	conditionCollection := client.Database("cleartoo").Collection("conditions")
	materialCollection := client.Database("cleartoo").Collection("materials")

	cartClient := cartPb.NewCartService("cleartoo.cart", service.Client())

	repository := &MongoRepository{productCollection, genderCollection, categoryCollection, sizeCollection, brandCollection, colorCollection, conditionCollection, materialCollection}

	h := &handler{repository, cartClient}

	if err := pb.RegisterCatalogServiceHandler(service.Server(), h); err != nil{
		fmt.Println(err)
	}

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}