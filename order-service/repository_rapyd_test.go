// orrder-service/main.go

package main

import(
	"context"
	"fmt"
	"errors"
	"net/http"
	"math/rand"
	_ "strings"
	"io"
	"strconv"
	_ "net/url"
	"bytes"
	"time"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
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
}

type Card struct{
	Number 			string
	ExpirationMonth int32
	ExpirationYear 	int32
	Cvv 			int32
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
		UserID: 		req.UserId,
	}
}

func MarshalCard(card *pb.Card) *Card{
	return &Card{
		Number:			card.Number,
		ExpirationMonth:card.ExpirationMonth,
		ExpirationYear:	card.ExpirationYear,
		Cvv:			card.Cvv,
	}
}

func UnmarshalCard(card *Card) *pb.Card{
	return &pb.Card{
		Number:			card.Number,
		ExpirationMonth:card.ExpirationMonth,
		ExpirationYear:	card.ExpirationYear,
		Cvv:			card.Cvv,
	}
}

func MarshalProduct(product *pb.Product) *Product{
	objId, _ := primitive.ObjectIDFromHex(product.Id)
	return &Product{
		ID:				objId,
		Available:		product.Available,
		Title:			product.Title,
		Price:			product.Price,
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

type CardPaymentRequest struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
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
}

type CardPaymentResponse struct {
	Status struct {
		ErrorCode    string `json:"error_code"`
		Status       string `json:"status"`
		Message      string `json:"message"`
		ResponseCode string `json:"response_code"`
		OperationID  string `json:"operation_id"`
	} `json:"status"`
	Data struct {
		ID                        string  `json:"id"`
		Amount                    float64 `json:"amount"`
		OriginalAmount            float64 `json:"original_amount"`
		IsPartial                 bool    `json:"is_partial"`
		CurrencyCode              string  `json:"currency_code"`
		CountryCode               string  `json:"country_code"`
		Status                    string  `json:"status"`
		CustomerToken             string  `json:"customer_token"`
		PaymentMethod             string  `json:"payment_method"`
		Expiration                int     `json:"expiration"`
		Captured                  bool    `json:"captured"`
		ErrorPaymentURL           string  `json:"error_payment_url"`
		CreatedAt                 int     `json:"created_at"`
		Paid                      bool    `json:"paid"`
		PaidAt                    int     `json:"paid_at"`
		PaymentMethodType         string  `json:"payment_method_type"`
		PaymentMethodTypeCategory string  `json:"payment_method_type_category"`
	} `json:"data"`
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
	CancelOrder(ctx context.Context, req *CancelOrderRequest) error
	AddTransaction(ctx context.Context, req *AddTransactionRequest) error
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

func (repo *MongoRepository) Order(ctx context.Context, req *OrderRequest) error{
	if req.Order.PaymentMethod == "card" {
		hc := http.Client{}
		paymentReq := CardPaymentRequest{
			Amount: strconv.Itoa(int(req.Order.Total)),
			Currency: "THB",
			PaymentMethod: PaymentMethod{
				Type: "th_visa_card",
				Fields: Fields{
					Number: "4111111111111111",
					ExpirationMonth: "10",
					ExpirationYear: "21",
					Cvv: "123",
				},
			},
			ErrorPaymentURL: "https://error_example.net",
			Capture: true,
		}
		jsonValue, err := json.Marshal(paymentReq)
		if err != nil {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("payment request marshal failed... %v", err))
		}
		httpReq, err := http.NewRequest("POST", "https://sandboxapi.rapyd.net/v1/payments", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("payment request creation failed... %v", err))
		}
		http_method := "post"
		url_path := "/v1/payments"
		salt := strconv.Itoa(10000000 + rand.Intn(99999999-10000000))
		timestamp := strconv.Itoa(int(time.Now().Unix()))
		body_string := form
		access_key := "CA818DDE1BDFE4CEDFFA"
		secret_key := "fea7f18fb788775d863870685b98078b9ab91e19ed13d5218b901b51968e1053955daaa7aad7b950"
		sign_str := http_method + url_path + salt + timestamp + access_key + secret_key + string(body_string[:])
	  	h := hmac.New(sha256.New, []byte(secret_key))
	  	h.Write([]byte(sign_str))
	  	buf := h.Sum(nil)
		signature := base64.RawURLEncoding.EncodeToString(buf)
		// req.PostForm = form
		httpReq.Header.Add("content-Type", "application/json")
		httpReq.Header.Add("access_key", access_key)
		httpReq.Header.Add("salt", salt)
		httpReq.Header.Add("signature", signature)
		httpReq.Header.Add("timestamp", timestamp)

		resp, err := hc.Do(httpReq)
		if err != nil {
			fmt.Println(err)
			return errors.New(fmt.Sprintf("payment request failed... %v", err))
		}
		data, err2 := io.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
			return errors.New(fmt.Sprintf("payment body lecture failed... %v", err2))
		}
		defer resp.Body.Close()
		fmt.Println(data)
		// if data.status != "SUCCESS" || data.data.status != "CLO" || data.data.amount != req.Order.Total {
		// 	return errors.New(fmt.Sprintf("payment failed... %v", data.status.message))
		// }
	}
	_, err := repo.orderCollection.InsertOne(ctx, req.Order)
	return err
}

func (repo *MongoRepository) GetOrders(ctx context.Context, req *GetRequest) ([]*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"userid", bson.D{bson.E{"$eq", req.UserID}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
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
	bsonFilters = append(bsonFilters, bson.E{"products.ownerid", bson.D{bson.E{"$elemMatch", req.UserID}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
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
	bsonFilters = append(bsonFilters, bson.E{"user_id", bson.D{bson.E{"$eq", req.UserID}}})
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
	_, err := repo.orderCollection.UpdateOne(ctx,bson.M{"_id": req.OrderID} , bson.M{"status": "confirmed"})
	return "confirmed", err
}

func (repo *MongoRepository) CancelOrder(ctx context.Context, req *CancelOrderRequest) error{
	_, err := repo.orderCollection.UpdateOne(ctx,bson.M{"_id": req.OrderID} , bson.M{"status": "canceled"})
	return err
}