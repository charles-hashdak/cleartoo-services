// catalog-service/main.go

package main

import(
	"fmt"
	"time"

	orderPb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	"github.com/micro/go-micro/v2"
)

func main(){

	service := micro.NewService(
		micro.Name("cleartoo.worker"),
	)

	service.Init()

	orderClient := orderPb.NewOrderService("cleartoo.order", service.Client())

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}

	for true {
        fmt.Println("Infinite Loop 2")
        checkInTransit(orderClient)
        time.Sleep(time.Second*300)
    }
}