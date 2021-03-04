package main

import (
	"context"
	"fmt"

	orderPb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
)

func checkInTransit(ctx context.Context, orderClient orderPb.OrderService) error {
	ordersRes, err := orderClient.GetInTransitOrders(ctx, &orderPb.GetRequest{})
	orders := ordersRes.Orders
	for _, order := range orders {
		fmt.Println("order")
		fmt.Println(order.Status)
	}
	return nil
	// call to orderPb to fetch sent orders
	// call to shippingPb for each orders to fetch thai post info
	// for delivered, actualise shipping status and order status
}