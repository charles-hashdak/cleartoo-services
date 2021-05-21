// orrder-service/main.go

package main

import(
	"context"
	"fmt"
	"errors"
	"net/http"
	_ "net/http/httputil"
	"math/rand"
	_ "strings"
	"io"
	"os"
	"strconv"
	_ "net/url"
	"bytes"
	"time"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
	_ "log"

	pb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID 			string 				`json:"user_id"`
	Products 		Products 			`json:"products"`
	SubTotal 		float32				`json:"sub_total"`
	ShippingFees    float32				`json:"shipping_fees"`
	Taxes 			float32				`json:"taxes"`
	Total 			float32				`json:"total"`
	Status 			string 				`json:"status"`
	ShippingMethod 	string 				`json:"shipping_method"`
	PaymentMethod 	string 				`json:"payment_method"`
	Address 		Address 			`json:"address"`
	TrackID 		string 				`json:"trackid"`
	ShippingStatus	string 				`json:"shippingstatus"`
}

type Address struct{
	Indications 	string 				`json:"indications"`
	AddressLine1 	string 				`json:"address_line1"`
	FirstName 		string 				`json:"first_name"`
	LastName 		string 				`json:"last_name"`
	Phone 			string 				`json:"phone"`
	Country 		string 				`json:"country"`
	City 			string 				`json:"city"`
	PostalCode		string 				`json:"postal_code"`
}

type Card struct{
	Number 			string
	ExpirationMonth string
	ExpirationYear 	string
	Cvv 			string
	HolderName		string
}

type Wallet struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID 			string 				`json:"user_id"`
	Balance 		int64 				`json:"balance"`
}

type Transaction struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	WalletID 		string 				`json:"wallet_id"`
	Amount 			int64 				`json:"amount"`
	Type 			string 				`json:"type"`
	OrderID 		string 				`json:"order_id"`
}

type Product struct{
	ID 				primitive.ObjectID
	Available 		bool
	Title 			string
	Price 			int32
	Weight 			int32
	Photo 			Photo
	Category 		string
	Size 			string
	Color1 			Color
	Color2 			Color
	Brand 			string
	Condition 		string
	Material 		string
	OwnerID 		string
}

type Photo struct{
	ID 				primitive.ObjectID 
	Url 			string
	IsMain 			bool
	Height 			int32
	Width 			int32
}

type Color struct{
	ID 				primitive.ObjectID
	Name 			string
	HexCode 		string
	Image 			string
}

type Products []*Product

type OrderRequest struct {
	UserID 			string
	Order 			Order
	Card 			Card
}

type OrderResponse struct {
	Added			bool
}

type DeleteFromOrderRequest struct {
	UserID 			string
	Product 		Product
}

type DeleteFromOrderResponse struct {
	Deleted			bool
}

type GetRequest struct {
	UserID 			string
}

type GetSingleRequest struct {
	OrderID 		string
}

type GetResponse struct {
	Order 			Order
}

type GetWalletRequest struct {
	UserID 		string
	WalletID 	string
}

type InitializeWalletRequest struct {
	UserID 			string
}

type UpdateWalletRequest struct {
	Wallet			Wallet
}

type AddTransactionRequest struct {
	Transaction		Transaction
}

type CancelOrderRequest struct {
	OrderID 		string
	UserID 			string
}

type CancelOrderResponse struct {
	Canceled 		bool
}

type UpdateOrderStatusRequest struct {
	OrderID 		string
	Status 			string
	TrackID 		string
	UserID 			string
}

type UpdateOrderStatusResponse struct {
	Updated 		bool
	Status 			string
}

func MarshalOrderRequest(req *pb.OrderRequest) *OrderRequest{
	return &OrderRequest{
		UserID: 		req.UserId,
		Order: 			*MarshalOrder(req.Order),
		Card: 			*MarshalCard(req.Card),
	}
}

func UnmarshalGetSingleRequest(req *GetSingleRequest) *pb.GetSingleRequest{
	return &pb.GetSingleRequest{
		OrderId: 		req.OrderID,
	}
}

func MarshalGetSingleRequest(req *pb.GetSingleRequest) *GetSingleRequest{
	return &GetSingleRequest{
		OrderID: 		req.OrderId,
	}
}

func UnmarshalOrderRequest(req *OrderRequest) *pb.OrderRequest{
	return &pb.OrderRequest{
		UserId: 		req.UserID,
		Order: 			UnmarshalOrder(&req.Order),
		Card: 			UnmarshalCard(&req.Card),
	}
}

func MarshalOrderResponse(req *pb.OrderResponse) *OrderResponse{
	return &OrderResponse{
		Added: 			req.Added,
	}
}

func UnmarshalOrderResponse(req *OrderResponse) *pb.OrderResponse{
	return &pb.OrderResponse{
		Added: 			req.Added,
	}
}

func MarshalDeleteFromOrderRequest(req *pb.DeleteFromOrderRequest) *DeleteFromOrderRequest{
	return &DeleteFromOrderRequest{
		UserID: 		req.UserId,
		Product: 		*MarshalProduct(req.Product),
	}
}

func UnmarshalDeleteFromOrderRequest(req *DeleteFromOrderRequest) *pb.DeleteFromOrderRequest{
	return &pb.DeleteFromOrderRequest{
		UserId: 		req.UserID,
		Product: 		UnmarshalProduct(&req.Product),
	}
}

func MarshalDeleteFromOrderResponse(req *pb.DeleteFromOrderResponse) *DeleteFromOrderResponse{
	return &DeleteFromOrderResponse{
		Deleted: 			req.Deleted,
	}
}

func UnmarshalDeleteFromOrderResponse(req *DeleteFromOrderResponse) *pb.DeleteFromOrderResponse{
	return &pb.DeleteFromOrderResponse{
		Deleted: 			req.Deleted,
	}
}

func MarshalGetRequest(req *pb.GetRequest) *GetRequest{
	return &GetRequest{
		UserID: 		req.UserId,
	}
}

func UnmarshalGetRequest(req *GetRequest) *pb.GetRequest{
	return &pb.GetRequest{
		UserId: 		req.UserID,
	}
}

func MarshalGetWalletRequest(req *pb.GetWalletRequest) *GetWalletRequest{
	return &GetWalletRequest{
		UserID: 		req.UserId,
		WalletID: 		req.WalletId,
	}
}

func MarshalInitializeWalletRequest(req *pb.InitializeWalletRequest) *InitializeWalletRequest{
	return &InitializeWalletRequest{
		UserID: 		req.UserId,
	}
}

func MarshalUpdateWalletRequest(req *pb.UpdateWalletRequest) *UpdateWalletRequest{
	return &UpdateWalletRequest{
		Wallet: 		*MarshalWallet(req.Wallet),
	}
}

func MarshalAddTransactionRequest(req *pb.AddTransactionRequest) *AddTransactionRequest{
	return &AddTransactionRequest{
		Transaction: 	*MarshalTransaction(req.Transaction),
	}
}

func MarshalCancelOrderRequest(req *pb.CancelOrderRequest) *CancelOrderRequest{
	return &CancelOrderRequest{
		OrderID: 		req.OrderId,
		UserID: 		req.UserId,
	}
}

func MarshalUpdateOrderStatusRequest(req *pb.UpdateOrderStatusRequest) *UpdateOrderStatusRequest{
	return &UpdateOrderStatusRequest{
		OrderID: 		req.OrderId,
		Status: 		req.Status,
		TrackID: 		req.TrackId,
		UserID: 		req.UserId,
	}
}

func MarshalCard(card *pb.Card) *Card{
	return &Card{
		Number:			card.Number,
		ExpirationMonth:card.ExpirationMonth,
		ExpirationYear:	card.ExpirationYear,
		Cvv:			card.Cvv,
		HolderName:		card.HolderName,
	}
}

func UnmarshalCard(card *Card) *pb.Card{
	return &pb.Card{
		Number:			card.Number,
		ExpirationMonth:card.ExpirationMonth,
		ExpirationYear:	card.ExpirationYear,
		Cvv:			card.Cvv,
		HolderName:		card.HolderName,
	}
}

func MarshalProduct(product *pb.Product) *Product{
	objId, _ := primitive.ObjectIDFromHex(product.Id)
	return &Product{
		ID:				objId,
		Available:		product.Available,
		Title:			product.Title,
		Price:			product.Price,
		Weight:			product.Weight,
		Photo:			*MarshalPhoto(product.Photo),
		Category:		product.Category,
		Size:			product.Size,
		Color1:			*MarshalColor(product.Color1),
		Color2:			*MarshalColor(product.Color2),
		Brand:			product.Brand,
		Condition:		product.Condition,
		Material:		product.Material,
		OwnerID:		product.OwnerId,
	}
}

func UnmarshalProduct(product *Product) *pb.Product{
	return &pb.Product{
		Id:				product.ID.Hex(),
		Available:		product.Available,
		Title:			product.Title,
		Price:			product.Price,
		Weight:			product.Weight,
		Photo:			UnmarshalPhoto(&product.Photo),
		Category:		product.Category,
		Size:			product.Size,
		Color1:			UnmarshalColor(&product.Color1),
		Color2:			UnmarshalColor(&product.Color2),
		Brand:			product.Brand,
		Condition:		product.Condition,
		Material:		product.Material,
		OwnerId:		product.OwnerID,
	}
}

func MarshalProducts(products []*pb.Product) Products {
	collection := make(Products, 0)
	for _, product := range products {
		collection = append(collection, MarshalProduct(product))
	}
	return collection
}

func UnmarshalProducts(products Products) []*pb.Product {
	collection := make([]*pb.Product, 0)
	for _, product := range products {
		collection = append(collection, UnmarshalProduct(product))
	}
	return collection
}

func MarshalWallet(wallet *pb.Wallet) *Wallet{
	objId, _ := primitive.ObjectIDFromHex(wallet.Id)
	return &Wallet{
		ID:				objId,
		UserID:			wallet.UserId,
		Balance:		wallet.Balance,
	}
}

func UnmarshalWallet(wallet *Wallet) *pb.Wallet{
	return &pb.Wallet{
		Id:				wallet.ID.Hex(),
		UserId:			wallet.UserID,
		Balance:		wallet.Balance,
	}
}

func MarshalTransaction(transaction *pb.Transaction) *Transaction{
	objId, _ := primitive.ObjectIDFromHex(transaction.Id)
	return &Transaction{
		ID:				objId,
		WalletID:		transaction.WalletId,
		Type:			transaction.Type,
		Amount:			transaction.Amount,
		OrderID:		transaction.OrderId,
	}
}

func UnmarshalTransaction(transaction *Transaction) *pb.Transaction{
	return &pb.Transaction{
		Id:				transaction.ID.Hex(),
		WalletId:		transaction.WalletID,
		Type:			transaction.Type,
		Amount:			transaction.Amount,
		OrderId:		transaction.OrderID,
	}
}

func MarshalOrder(order *pb.Order) *Order{
	objId, _ := primitive.ObjectIDFromHex(order.Id)
	return &Order{
		ID:				objId,
		UserID:			order.UserId,
		Products:		MarshalProducts(order.Products),
		SubTotal:		order.SubTotal,
		ShippingFees:	order.ShippingFees,
		Taxes:			order.Taxes,
		Total:			order.Total,
		Status:			order.Status,
		ShippingMethod:	order.ShippingMethod,
		PaymentMethod:	order.PaymentMethod,
		Address:		*MarshalAddress(order.Address),
		TrackID:		order.TrackId,
		ShippingStatus:	order.ShippingStatus,
	}
}

func UnmarshalOrder(order *Order) *pb.Order{
	return &pb.Order{
		Id:				order.ID.Hex(),
		UserId:			order.UserID,
		Products:		UnmarshalProducts(order.Products),
		SubTotal:		order.SubTotal,
		ShippingFees:	order.ShippingFees,
		Taxes:			order.Taxes,
		Total:			order.Total,
		Status:			order.Status,
		ShippingMethod:	order.ShippingMethod,
		PaymentMethod:	order.PaymentMethod,
		Address:		UnmarshalAddress(&order.Address),
		ShippingStatus:	order.ShippingStatus,
	}
}

func MarshalAddress(address *pb.Address) *Address{
	return &Address{
		Indications: 	address.Indications,
		AddressLine1: 	address.AddressLine1,
		FirstName: 		address.FirstName,
		LastName: 		address.LastName,
		Phone: 			address.Phone,
		Country: 		address.Country,
		City: 			address.City,
		PostalCode:		address.PostalCode,
	}
}

func UnmarshalAddress(address *Address) *pb.Address{
	return &pb.Address{
		Indications: 	address.Indications,
		AddressLine1: 	address.AddressLine1,
		FirstName: 		address.FirstName,
		LastName: 		address.LastName,
		Phone: 			address.Phone,
		Country: 		address.Country,
		City: 			address.City,
		PostalCode:		address.PostalCode,
	}
}

func MarshalColor(color *pb.Color) *Color{
	if(color == nil){
		return &Color{}
	}
	objId, _ := primitive.ObjectIDFromHex(color.Id)
	return &Color{
		ID:				objId,
		Name:			color.Name,
		HexCode:		color.HexCode,
		Image:			color.Image,
	}
}

func UnmarshalColor(color *Color) *pb.Color{
	if(color == nil){
		return &pb.Color{}
	}
	return &pb.Color{
		Id:				color.ID.Hex(),
		Name:			color.Name,
		HexCode:		color.HexCode,
		Image:			color.Image,
	}
}

func MarshalPhoto(photo *pb.Photo) *Photo{
	if(photo == nil){
		return &Photo{}
	}
	objId, _ := primitive.ObjectIDFromHex(photo.Id)
	return &Photo{
		ID:				objId,
		Url: 			photo.Url,
		IsMain:			photo.IsMain,
		Height:			photo.Height,
		Width:			photo.Width,
	}
}

func UnmarshalPhoto(photo *Photo) *pb.Photo{
	if(photo == nil){
		return &pb.Photo{}
	}
	return &pb.Photo{
		Id:				photo.ID.Hex(),
		Url: 			photo.Url,
		IsMain:			photo.IsMain,
		Height:			photo.Height,
		Width:			photo.Width,
	}
}

func UnmarshalOrderCollection(orders []*Order) []*pb.Order {
	collection := make([]*pb.Order, 0)
	for _, order := range orders {
		collection = append(collection, UnmarshalOrder(order))
	}
	return collection
}

type repository interface{
	Order(ctx context.Context, req *OrderRequest) error
	GetSingleOrder(ctx context.Context, req *GetSingleRequest) (*Order, error)
	GetOrders(ctx context.Context, req *GetRequest) ([]*Order, error)
	GetInTransitOrders(ctx context.Context, req *GetRequest) ([]*Order, error)
	GetSales(ctx context.Context, req *GetRequest) ([]*Order, error)
	GetWallet(ctx context.Context, req *GetWalletRequest) (*Wallet, error)
	InitializeWallet(ctx context.Context, req *InitializeWalletRequest) error
	UpdateWallet(ctx context.Context, req *UpdateWalletRequest) error
	UpdateOrderStatus(ctx context.Context, req *UpdateOrderStatusRequest) (string, error)
	UpdateOrderShippingStatus(ctx context.Context, req *UpdateOrderStatusRequest) (string, error)
	CancelOrder(ctx context.Context, req *CancelOrderRequest) error
	AddTransaction(ctx context.Context, req *AddTransactionRequest) error
	GetThaiPostToken(ctx context.Context, req *UpdateOrderStatusRequest) (string, error)
	GetThaiPostStatus(ctx context.Context, req *UpdateOrderStatusRequest) (bool, error)
}

type MongoRepository struct{
	orderCollection 		*mongo.Collection
	walletCollection 		*mongo.Collection
	transactionCollection 	*mongo.Collection
}

func (repo *MongoRepository) GetSingleOrder(ctx context.Context, req *GetSingleRequest) (*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"_id", bson.D{bson.E{"$eq", req.OrderID}}})
	//bsonFilters = append(bsonFilters, bson.E{"available", bson.D{bson.E{"$eq", true}}})
	var order *Order
	err := repo.orderCollection.FindOne(ctx, bsonFilters, nil).Decode(&order)
	return order, err
}

type CardRequest struct {
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	ErrorPaymentURL string `json:"error_payment_url"`
	Capture         bool   `json:"capture"`
}

type PaymentMethod struct {
	Type   string `json:"type"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Number          string `json:"number"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	Cvv             string `json:"cvv"`
	Name            string `json:"name"`
}

func (repo *MongoRepository) Order(ctx context.Context, req *OrderRequest) error{
	fmt.Println("ok")
	if req.Order.PaymentMethod == "card" {
		hc := http.Client{}
		form, _ := json.Marshal(CardRequest{
		   Amount: req.Order.Total,
		   Currency: "THB",
		   PaymentMethod: PaymentMethod{
		       Type: "th_visa_card",
		       Fields: Fields{
		           Number: req.Card.Number,
		           ExpirationMonth: req.Card.ExpirationMonth,
		           ExpirationYear: req.Card.ExpirationYear,
		           Cvv: req.Card.Cvv,
		           Name: req.Card.HolderName,
		       },
		   },
		  ErrorPaymentURL: "https://error_example.net",
		  Capture: true})
		http_method := "post"
		url_path := "/v1/payments"
		salt := strconv.Itoa(10000000 + rand.Intn(99999999-10000000))
		timestamp := strconv.Itoa(int(time.Now().Unix()))
		access_key := os.Getenv("RAPYD_ACCESS_KEY")
		secret_key := os.Getenv("RAPYD_SECRET_KEY")
		sign_str := http_method + url_path + salt + timestamp + access_key + secret_key + string(form[:])
	  	h := hmac.New(sha256.New, []byte(secret_key))
	  	h.Write([]byte(sign_str))
	  	buf := h.Sum(nil)
	  	hex := hex.EncodeToString(buf)
		signature := base64.URLEncoding.EncodeToString([]byte(hex))
		httpReq, err := http.NewRequest("POST", "https://sandboxapi.rapyd.net/v1/payments", bytes.NewBuffer(form))
		// req.PostForm = form
		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Header.Set("access_key", access_key)
		httpReq.Header.Set("salt", salt)
		httpReq.Header.Set("signature", signature)
		httpReq.Header.Set("timestamp", timestamp)

		resp, err := hc.Do(httpReq)
		if err != nil {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("payment request failed... %v", err))
		}
		if err != nil {
			fmt.Println(err)
		}

		data, err2 := io.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
			return errors.New(fmt.Sprintf("payment body lecture failed... %v", err2))
		}
		var result map[string]interface{}
		fmt.Println(string(data))
		json.Unmarshal([]byte(string(data)), &result)
		status := result["status"].(map[string]interface{})
		if status["status"] != "SUCCESS" {
			return errors.New(fmt.Sprintf("payment failed... %v", status["message"]))
		}else{
			// resultData := result["data"].(map[string]interface{})
			// if resultData["status"] != "CLO" || resultData["amount"] != req.Order.Total {
			// 	return errors.New(fmt.Sprintf("payment failed... %v", status["message"]))
			// }
		}
		resp.Body.Close()
	}
	_, err := repo.orderCollection.InsertOne(ctx, req.Order)
	return err
}

func (repo *MongoRepository) GetOrders(ctx context.Context, req *GetRequest) ([]*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"userid", bson.D{bson.E{"$eq", req.UserID}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var orders []*Order
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func (repo *MongoRepository) GetInTransitOrders(ctx context.Context, req *GetRequest) ([]*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"status", bson.D{bson.E{"$eq", "sent"}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var orders []*Order
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func (repo *MongoRepository) GetSales(ctx context.Context, req *GetRequest) ([]*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"products", bson.D{{"$elemMatch", bson.D{{"ownerid", req.UserID}}}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var orders []*Order
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func (repo *MongoRepository) GetWallet(ctx context.Context, req *GetWalletRequest) (*Wallet, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"userid", bson.D{bson.E{"$eq", req.UserID}}})
	var wallet *Wallet
	err := repo.walletCollection.FindOne(ctx, bsonFilters, nil).Decode(&wallet)
	return wallet, err
}

func (repo *MongoRepository) InitializeWallet(ctx context.Context, req *InitializeWalletRequest) error{
	wallet := &Wallet{
		UserID: req.UserID,
		Balance: 0,
	}
	_, err := repo.walletCollection.InsertOne(ctx, wallet)
	return err
}

func (repo *MongoRepository) AddTransaction(ctx context.Context, req *AddTransactionRequest) error{
	_, err := repo.transactionCollection.InsertOne(ctx, req.Transaction)
	return err
}

func (repo *MongoRepository) UpdateWallet(ctx context.Context, req *UpdateWalletRequest) error{
	_, err := repo.walletCollection.UpdateOne(ctx,bson.M{"_id": req.Wallet.ID} , req.Wallet)
	return err
}

func (repo *MongoRepository) UpdateOrderStatus(ctx context.Context, req *UpdateOrderStatusRequest) (string, error){
	orderId, _ := primitive.ObjectIDFromHex(req.OrderID)
	if req.Status == "sent" {
		status, err := repo.GetThaiPostStatus(ctx, req)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		if status == false {
			return "", errors.New(fmt.Sprintf("no package found with this id"))
		}
		update := bson.M{
		    "$set": bson.M{
		      "status": req.Status,
		      "trackid": req.TrackID,
		    },
	  	}
		_, err = repo.orderCollection.UpdateOne(ctx, bson.M{"_id": orderId}, update)
		return req.Status, err
	}else{
		update := bson.M{
		    "$set": bson.M{
		      "status": req.Status,
		    },
	  	}
		_, err := repo.orderCollection.UpdateOne(ctx, bson.M{"_id": orderId}, update)
		return req.Status, err
	}
}

func (repo *MongoRepository) UpdateOrderShippingStatus(ctx context.Context, req *UpdateOrderStatusRequest) (string, error){
	orderId, _ := primitive.ObjectIDFromHex(req.OrderID)
	update := bson.M{
	    "$set": bson.M{
	      "shippingstatus": req.Status,
	    },
  	}
	_, err := repo.orderCollection.UpdateOne(ctx, bson.M{"_id": orderId}, update)
	return req.Status, err
}

func (repo *MongoRepository) CancelOrder(ctx context.Context, req *CancelOrderRequest) error{
	_, err := repo.orderCollection.UpdateOne(ctx,bson.M{"_id": req.OrderID} , bson.M{"status": "canceled"})
	return err
}

func (repo *MongoRepository) GetThaiPostToken(ctx context.Context, req *UpdateOrderStatusRequest) (string, error){
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

func (repo *MongoRepository) GetThaiPostStatus(ctx context.Context, req *UpdateOrderStatusRequest) (bool, error){
	hc := http.Client{}
	// token := os.Getenv("THAI_POST_TOKEN")
	form, _ := json.Marshal(GetThaiPostStatusRequest{
		Status: "all",
		Language: "EN",
		Barcode: []string{req.TrackID},
	})
	httpReq, err := http.NewRequest("POST", "https://trackapi.thailandpost.co.th/post/api/v1/track", bytes.NewBuffer(form))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Token eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzZWN1cmUtYXBpIiwiYXVkIjoic2VjdXJlLWFwcCIsInN1YiI6IkF1dGhvcml6YXRpb24iLCJleHAiOjE2MjM2OTM0MzAsInJvbCI6WyJST0xFX1VTRVIiXSwiZCpzaWciOnsicCI6InpXNzB4IiwicyI6bnVsbCwidSI6IjhmYWM3Yjc3MWZiYTFjYjQ2ZGFhZmQ2NDE4NDdkM2JhIiwiZiI6InhzeiM5In19.Ql6_iQo8EhAl3SCzNW6PAmpamNwTNgdFn_gppAFxZXiXQInvYnonsqLxnEr2gR5VtkXVRF7lyiPdOXlfpG-NfQ")

	resp, err := hc.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return false, errors.New(fmt.Sprintf("thai post request failed... %v", err))
	}

	data, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return false, errors.New(fmt.Sprintf("pthai post body lecture failed... %v", err2))
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
		item, _ := json.Marshal(itemsResult[req.TrackID])
		if string(item) == "[]" {
			return false, errors.New(fmt.Sprintf("no package found..."))
		}
		return true, nil
	}
	resp.Body.Close()
	return false, nil
}