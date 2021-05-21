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
		ord, _ := json.Marshal(order)
		fmt.Println(string(ord))
		// statusCode, err := GetThaiPostStatus(order.TrackId)
		// fmt.Println(statusCode)
	}
	return nil
	// call to orderPb to fetch sent orders
	// call to shippingPb for each orders to fetch thai post info
	// for delivered, actualise shipping status and order status
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
		return "", errors.New(fmt.Sprintf("pthai post body lecture failed... %v", err2))
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

func GetThaiPostStatus(trackId string) (string, error){
	hc := http.Client{}
	// token := os.Getenv("THAI_POST_TOKEN")
	form, _ := json.Marshal(GetThaiPostStatusRequest{
		Status: "all",
		Language: "EN",
		Barcode: []string{trackId},
	})
	httpReq, err := http.NewRequest("POST", "https://trackapi.thailandpost.co.th/post/api/v1/track", bytes.NewBuffer(form))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Token eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzZWN1cmUtYXBpIiwiYXVkIjoic2VjdXJlLWFwcCIsInN1YiI6IkF1dGhvcml6YXRpb24iLCJleHAiOjE2MjM2OTM0MzAsInJvbCI6WyJST0xFX1VTRVIiXSwiZCpzaWciOnsicCI6InpXNzB4IiwicyI6bnVsbCwidSI6IjhmYWM3Yjc3MWZiYTFjYjQ2ZGFhZmQ2NDE4NDdkM2JhIiwiZiI6InhzeiM5In19.Ql6_iQo8EhAl3SCzNW6PAmpamNwTNgdFn_gppAFxZXiXQInvYnonsqLxnEr2gR5VtkXVRF7lyiPdOXlfpG-NfQ")

	resp, err := hc.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(fmt.Sprintf("thai post request failed... %v", err))
	}

	data, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return "", errors.New(fmt.Sprintf("pthai post body lecture failed... %v", err2))
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
			return "", errors.New(fmt.Sprintf("no package found..."))
		}
		var itemResult []interface{}
		json.Unmarshal([]byte(string(item)), &itemResult)
		lastStatus, _ := json.Marshal(itemResult[len(itemResult)-1])
		var lastStatusResult map[string]interface{}
		json.Unmarshal([]byte(string(lastStatus)), &lastStatusResult)
		lastStatusCode, _ := json.Marshal(lastStatusResult["code"])
		return string(lastStatusCode), nil
	}
	resp.Body.Close()
	return "", nil
}