package main

import (
	"context"
	"fmt"
	"errors"
	"net/http"
	"io"
	"os"
    "encoding/json"
	"bytes"

	orderPb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
)

func checkInTransit(orderClient orderPb.OrderService) error {
	ordersRes, err := orderClient.GetInTransitOrders(context.Background(), &orderPb.GetRequest{})
	if err != nil {
		return err
	}
	orders := ordersRes.Orders
	for _, order := range orders {
		statusCode, statusDescription, err := GetThaiPostStatus(order.TrackId)
		if err != nil {
			fmt.Println(err)
			// return errors.New(fmt.Sprintf("status request failed... %v", err))
		}
		updateOrderShippingStatusRes, err := orderClient.UpdateOrderShippingStatus(context.Background(), &orderPb.UpdateOrderStatusRequest{
			OrderId: order.Id,
			Status: statusDescription,
		})
		if err != nil || updateOrderShippingStatusRes.Updated == false {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("update shipping status request failed... %v", err))
		}
		if statusCode == "\"501\"" {
			updateOrderStatusRes, err := orderClient.UpdateOrderStatus(context.Background(), &orderPb.UpdateOrderStatusRequest{
				OrderId: order.Id,
				Status: "delivered",
			})
			if err != nil || updateOrderStatusRes.Updated == false {
				fmt.Println(err)
				return errors.New(fmt.Sprintf("update status request failed... %v", err))
			}
		}
	}
	return nil
}

func checkDelivered(orderClient orderPb.OrderService) error {
	ordersRes, err := orderClient.GetOrdersByStatus(context.Background(), &orderPb.GetRequest{
		Status: "delivered",
	})
	if err != nil {
		return err
	}
	orders := ordersRes.Orders
	for _, order := range orders {
		layout := "2006-01-02 15:04:05"
		orderUpdatedAt, _ := time.Parse(layout, order.UpdatedAt)
		if time.Since(orderUpdatedAt).Hours() >= 72 {
			updateOrderStatusRes, err := orderClient.UpdateOrderStatus(context.Background(), &orderPb.UpdateOrderStatusRequest{
				OrderId: order.Id,
				Status: "finalised",
			})
			if err != nil || updateOrderStatusRes.Updated == false {
				fmt.Println(err)
				return errors.New(fmt.Sprintf("update status request failed... %v", err))
			}
		}
	}
	return nil
}

func checkConfirmed(orderClient orderPb.OrderService) error {
	ordersRes, err := orderClient.GetOrdersByStatus(context.Background(), &orderPb.GetRequest{
		Status: "confirmed",
	})
	if err != nil {
		return err
	}
	orders := ordersRes.Orders
	for _, order := range orders {
		layout := "2006-01-02 15:04:05"
		orderUpdatedAt, _ := time.Parse(layout, order.UpdatedAt)
		if time.Since(orderUpdatedAt).Hours() >= 120 {
			updateOrderStatusRes, err := orderClient.UpdateOrderStatus(context.Background(), &orderPb.UpdateOrderStatusRequest{
				OrderId: order.Id,
				Status: "canceled",
			})
			if err != nil || updateOrderStatusRes.Updated == false {
				fmt.Println(err)
				return errors.New(fmt.Sprintf("update status request failed... %v", err))
			}
		}
	}
	return nil
}

func checkWaiting(orderClient orderPb.OrderService) error {
	ordersRes, err := orderClient.GetOrdersByStatus(context.Background(), &orderPb.GetRequest{
		Status: "waiting",
	})
	if err != nil {
		return err
	}
	orders := ordersRes.Orders
	for _, order := range orders {
		layout := "2006-01-02 15:04:05"
		orderUpdatedAt, _ := time.Parse(layout, order.UpdatedAt)
		if time.Since(orderUpdatedAt).Hours() >= 96 {
			updateOrderStatusRes, err := orderClient.UpdateOrderStatus(context.Background(), &orderPb.UpdateOrderStatusRequest{
				OrderId: order.Id,
				Status: "canceled",
			})
			if err != nil || updateOrderStatusRes.Updated == false {
				fmt.Println(err)
				return errors.New(fmt.Sprintf("update status request failed... %v", err))
			}
		}
	}
	return nil
}

func GetThaiPostToken() (string, error){
	hc := http.Client{}
	api_token := os.Getenv("THAI_POST_TOKEN")
	httpReq, err := http.NewRequest("POST", "https://trackapi.thailandpost.co.th/post/api/v1/authenticate/token", nil)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Token "+api_token)

	resp, err := hc.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(fmt.Sprintf("thai post request failed... %v", err))
	}

	data, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return "", errors.New(fmt.Sprintf("thai post body lecture failed... %v", err2))
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(data)), &result)
	thai_post_token, _ := json.Marshal(result["token"])
	resp.Body.Close()
	return string(thai_post_token), nil
}

type GetThaiPostStatusRequest struct {
	Status   string   `json:"status"`
	Language string   `json:"language"`
	Barcode  []string `json:"barcode"`
}

func GetThaiPostStatus(trackId string) (string, string, error){
	hc := http.Client{}
	token, err := GetThaiPostToken()
	form, _ := json.Marshal(GetThaiPostStatusRequest{
		Status: "all",
		Language: "EN",
		Barcode: []string{trackId},
	})
	httpReq, err := http.NewRequest("POST", "https://trackapi.thailandpost.co.th/post/api/v1/track", bytes.NewBuffer(form))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Token "+token[1:len(token)-1])

	resp, err := hc.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return "", "", errors.New(fmt.Sprintf("thai post request failed... %v", err))
	}

	data, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return "", "", errors.New(fmt.Sprintf("pthai post body lecture failed... %v", err2))
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(data)), &result)
	status, _ := json.Marshal(result["status"])
	if string(status) == "true" {
		response, _ := json.Marshal(result["response"])
		var responseResult map[string]interface{}
		json.Unmarshal([]byte(string(response)), &responseResult)
		items, _ := json.Marshal(responseResult["items"])
		var itemsResult map[string]interface{}
		json.Unmarshal([]byte(string(items)), &itemsResult)
		item, _ := json.Marshal(itemsResult[trackId])
		if string(item) == "[]" {
			return "", "", errors.New(fmt.Sprintf("no package found for id "+trackId+"..."))
		}
		var itemResult []interface{}
		json.Unmarshal([]byte(string(item)), &itemResult)
		lastStatus, _ := json.Marshal(itemResult[len(itemResult)-1])
		var lastStatusResult map[string]interface{}
		json.Unmarshal([]byte(string(lastStatus)), &lastStatusResult)
		lastStatusCode, _ := json.Marshal(lastStatusResult["status"])
		lastStatusDescription, _ := json.Marshal(lastStatusResult["status_description"])
		return string(lastStatusCode), string(lastStatusDescription)[1:len(string(lastStatusDescription))-1], nil
	}
	resp.Body.Close()
	return "", "", nil
}