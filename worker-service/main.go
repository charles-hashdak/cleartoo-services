// catalog-service/main.go

package main

import(
	"fmt"
	"time"

	//cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	"github.com/micro/go-micro/v2"
)

func main(){

	service := micro.NewService(
		micro.Name("cleartoo.worker"),
	)

	service.Init()

	//cartClient := cartPb.NewCartService("cleartoo.cart", service.Client())

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}

	for true {
        fmt.Println("Infinite Loop 2")
        time.Sleep(time.Second*300)
    }
}