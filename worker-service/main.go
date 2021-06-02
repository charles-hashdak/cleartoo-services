// worker-service/main.go

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

	for true {
        checkInTransit(orderClient)
        checkDelivered(orderClient)
        checkConfirmed(orderClient)
        checkWaiting(orderClient)
        time.Sleep(time.Second*60)
    }

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}