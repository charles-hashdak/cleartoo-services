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
	Offers 			Offers 				`json:"offers"`
	CreatedAt 		string 				`json:"created_at"`
	UpdatedAt 		string 				`json:"updated_at"`
}

type Offer struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Amount			float32 			`json:"amount"`
	Status 			string 				`json:"status"`
	CreatedAt 		string 				`json:"created_at"`
	UpdatedAt 		string 				`json:"updated_at"`
}

type Offers []*Offer

type CreateOfferRequest struct{
	OrderID 		string
	Offer 			Offer
}

type WithdrawRequest struct{
	UserID 			string
	Amount 			float32
	BankAccount 	string
	BankName 		string
}

type Address struct{
	Title 			string 				`json:"title"`
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
	Balance 		float32 			`json:"balance"`
}

type Transaction struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	WalletID 		string 				`json:"wallet_id"`
	Amount 			float32 			`json:"amount"`
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
		Offers:			MarshalOffers(order.Offers),
		CreatedAt:		order.CreatedAt,
		UpdatedAt:		order.UpdatedAt,
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
		TrackId:		order.TrackID,
		ShippingStatus:	order.ShippingStatus,
		Offers:			UnmarshalOffers(order.Offers),
		CreatedAt:		order.CreatedAt,
		UpdatedAt:		order.UpdatedAt,
	}
}

func MarshalOffer(offer *pb.Offer) *Offer{
	if(offer == nil){
		return &Offer{}
	}
	objId, _ := primitive.ObjectIDFromHex(offer.Id)
	return &Offer{
		ID:				objId,
		Amount:			offer.Amount,
		Status:			offer.Status,
		CreatedAt:		offer.CreatedAt,
		UpdatedAt:		offer.UpdatedAt,
	}
}

func MarshalOffers(offers []*pb.Offer) Offers {
	collection := make(Offers, 0)
	for _, offer := range offers {
		collection = append(collection, MarshalOffer(offer))
	}
	return collection
}

func UnmarshalOffer(offer *Offer) *pb.Offer{
	if(offer == nil){
		return &pb.Offer{}
	}
	return &pb.Offer{
		Id:				offer.ID.Hex(),
		Amount:			offer.Amount,
		Status:			offer.Status,
		CreatedAt:		offer.CreatedAt,
		UpdatedAt:		offer.UpdatedAt,
	}
}

func UnmarshalOffers(offers Offers) []*pb.Offer {
	collection := make([]*pb.Offer, 0)
	for _, offer := range offers {
		collection = append(collection, UnmarshalOffer(offer))
	}
	return collection
}

func MarshalCreateOfferRequest(req *pb.CreateOfferRequest) *CreateOfferRequest{
	return &CreateOfferRequest{
		OrderID: 		req.OrderId,
		Offer: 			*MarshalOffer(req.Offer),
	}
}

func MarshalWithdrawRequest(req *pb.WithdrawRequest) *WithdrawRequest{
	return &WithdrawRequest{
		UserID: 		req.UserId,
		Amount: 		req.Amount,
		BankAccount: 	req.BankAccount,
		BankName: 		req.BankName,
	}
}

func MarshalAddress(address *pb.Address) *Address{
	return &Address{
		Title: 			address.Title,
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
		Title: 			address.Title,
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
	Order(ctx context.Context, req *OrderRequest) (string, error)
	GetSingleOrder(ctx context.Context, req *GetSingleRequest) (*Order, error)
	GetOrders(ctx context.Context, req *GetRequest) ([]*Order, error)
	GetOrder(ctx context.Context, orderId primitive.ObjectID) (*Order, error)
	GetOrderByOfferId(ctx context.Context, offerId primitive.ObjectID) (*Order, error)
	GetInTransitOrders(ctx context.Context, req *GetRequest) ([]*Order, error)
	GetSales(ctx context.Context, req *GetRequest) ([]*Order, error)
	CreateOffer(ctx context.Context, offer *CreateOfferRequest) error
	Withdraw(ctx context.Context, req *WithdrawRequest) error
	EditOffer(ctx context.Context, offer *Offer) error
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
	orderId, _ := primitive.ObjectIDFromHex(req.OrderID)
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"_id", bson.D{bson.E{"$eq", orderId}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		return order, nil
	}
	return nil, nil
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

func (repo *MongoRepository) Order(ctx context.Context, req *OrderRequest) (string, error){
	req.Order.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	req.Order.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	req.Order.ID = primitive.NewObjectID()
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
			return "", errors.New(fmt.Sprintf("payment request failed... %v", err))
		}

		data, err2 := io.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
			return "", errors.New(fmt.Sprintf("payment body lecture failed... %v", err2))
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(string(data)), &result)
		status := result["status"].(map[string]interface{})
		if status["status"] != "SUCCESS" {
			return "", errors.New(fmt.Sprintf("payment failed... %v", status["message"]))
		}else{
			// resultData := result["data"].(map[string]interface{})
			// if resultData["status"] != "CLO" || resultData["amount"] != req.Order.Total {
			// 	return "", errors.New(fmt.Sprintf("payment failed... %v", status["message"]))
			// }
		}
		resp.Body.Close()
	}else if req.Order.PaymentMethod == "wallet" {
		wallet, err := repo.GetWallet(
			ctx,
			&GetWalletRequest{
				UserID: req.Order.UserID,
			},
		)
		if err != nil {
			return "", errors.New(fmt.Sprintf("get wallet request failed... %v", err))
		}
		balance := wallet.Balance
		if balance < req.Order.Total {
			return "", errors.New(fmt.Sprintf("insuffisant balance..."))
		}else{
			err = repo.AddTransaction(
				ctx,
				&AddTransactionRequest{
					Transaction: Transaction{
						WalletID: wallet.ID.Hex(),
						Amount: req.Order.Total,
						Type: "payment",
						OrderID: req.Order.ID.Hex(),
					},
				},
			)
			if err != nil {
				return "", errors.New(fmt.Sprintf("add transaction request failed... %v", err))
			}
		}
	}
	insertResult, err := repo.orderCollection.InsertOne(ctx, req.Order)
	id, _ := insertResult.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
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

func (repo *MongoRepository) GetOrder(ctx context.Context, orderId primitive.ObjectID) (*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"_id", bson.D{bson.E{"$eq", orderId}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		return order, nil
	}
	return nil, nil
}

func (repo *MongoRepository) GetOrderByOfferId(ctx context.Context, offerId primitive.ObjectID) (*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"offers", bson.D{{"$elemMatch", bson.D{{"_id", offerId}}}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		return order, nil
	}
	return nil, nil
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

func (repo *MongoRepository) CreateOffer(ctx context.Context, req *CreateOfferRequest) error{
	req.Offer.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	req.Offer.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	req.Offer.ID = primitive.NewObjectID()
	req.Offer.Status = "pending"
	orderId, _ := primitive.ObjectIDFromHex(req.OrderID)
	order, err := repo.GetOrder(ctx, orderId)
	if req.Offer.Amount > order.SubTotal + order.ShippingFees {
		return errors.New(fmt.Sprintf("reimbursment amount can't exceed order subtotal and shipping fees..."))
	}
	_, err = repo.orderCollection.UpdateOne(
	    ctx,
	    bson.M{"_id": orderId},
	    bson.D{
	        {"$push", bson.D{{"offers", req.Offer}}},
	    },
	)
	return err
}

func (repo *MongoRepository) Withdraw(ctx context.Context, req *WithdrawRequest) error{
	wallet, err := repo.GetWallet(
		ctx,
		&GetWalletRequest{
			UserID: req.UserID,
		},
	)
	if err != nil {
		return errors.New(fmt.Sprintf("get wallet request failed... %v", err))
	}
	balance := wallet.Balance
	if balance < req.Amount {
		return errors.New(fmt.Sprintf("insuffisant balance..."))
	}else{
		err = repo.AddTransaction(
			ctx,
			&AddTransactionRequest{
				Transaction: Transaction{
					WalletID: wallet.ID.Hex(),
					Amount: req.Amount,
					Type: "withdrawal",
				},
			},
		)
		if err != nil {
			return errors.New(fmt.Sprintf("add transaction request failed... %v", err))
		}
	}
	return nil
}

func (repo *MongoRepository) EditOffer(ctx context.Context, offer *Offer) error{
	offer.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := repo.orderCollection.UpdateOne(
	    ctx,
	    bson.M{"offers._id": offer.ID},
	    bson.M{"$set": bson.M{"offers.$.status": offer.Status, "offers.$.updatedat": offer.UpdatedAt}},
	)
	if offer.Status == "accepted" {
		order, err := repo.GetOrderByOfferId(ctx, offer.ID)
		if err != nil {
			return errors.New(fmt.Sprintf("get order request failed... %v", err))
		}
		wallet, err := repo.GetWallet(
			ctx,
			&GetWalletRequest{
				UserID: order.UserID,
			},
		)
		if err != nil {
			return errors.New(fmt.Sprintf("get wallet request failed... %v", err))
		}
		err = repo.AddTransaction(
			ctx,
			&AddTransactionRequest{
				Transaction: Transaction{
					WalletID: wallet.ID.Hex(),
					Amount: offer.Amount,
					Type: "partial_cancel",
					OrderID: order.ID.Hex(),
				},
			},
		)
		if err != nil {
			return errors.New(fmt.Sprintf("add transaction request failed... %v", err))
		}
		status := "finalised"
		if offer.Amount == (order.SubTotal + order.ShippingFees) {
			status = "returned"
		}
		_, err = repo.UpdateOrderStatus(
			ctx,
			&UpdateOrderStatusRequest{
				OrderID: order.ID.Hex(),
				Status: status,
			},
		)
		if err != nil {
			return errors.New(fmt.Sprintf("add transaction request failed... %v", err))
		}
	}
	return err
}

func (repo *MongoRepository) GetWallet(ctx context.Context, req *GetWalletRequest) (*Wallet, error){
	walletId, _ := primitive.ObjectIDFromHex(req.WalletID)
	bsonFilters := bson.M{"$or": []bson.M{bson.M{"userid": req.UserID}, bson.M{"_id": walletId}}}
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.walletCollection.Find(ctx,  bsonFilters, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cur.Next(ctx) {
		var wallet *Wallet
		if err := cur.Decode(&wallet); err != nil {
			return nil, err
		}
		return wallet, nil
	}
	return nil, errors.New(fmt.Sprintf("no wallet found..."))
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
	if err != nil {
		return errors.New(fmt.Sprintf("insert transaction failed... %v", err))
	}
	wallet, err := repo.GetWallet(
		ctx,
		&GetWalletRequest{
			WalletID: req.Transaction.WalletID,
		},
	)
	if err != nil {
		return errors.New(fmt.Sprintf("get wallet request failed... %v", err))
	}
	if req.Transaction.Type == "payment" || req.Transaction.Type == "withdrawal" {
		wallet.Balance = wallet.Balance - req.Transaction.Amount
		err = repo.UpdateWallet(
			ctx,
			&UpdateWalletRequest{
				Wallet: *wallet,
			},
		)
		if err != nil {
			return errors.New(fmt.Sprintf("update wallet request failed... %v", err))
		}
		if req.Transaction.Type == "withdrawal" {
			//send mail
		}
	}
	if req.Transaction.Type == "gain" || req.Transaction.Type == "cancel" || req.Transaction.Type == "partial_cancel" {
		wallet.Balance = wallet.Balance + req.Transaction.Amount
		err = repo.UpdateWallet(
			ctx,
			&UpdateWalletRequest{
				Wallet: *wallet,
			},
		)
		if err != nil {
			return errors.New(fmt.Sprintf("update wallet request failed... %v", err))
		}
	}
	return nil
}

func (repo *MongoRepository) UpdateWallet(ctx context.Context, req *UpdateWalletRequest) error{
	update := bson.M{
	    "$set": bson.M{
	      "balance": req.Wallet.Balance,
	    },
	  }
	_, err := repo.walletCollection.UpdateOne(ctx, bson.M{"_id": req.Wallet.ID}, update)
	return err
}

func (repo *MongoRepository) UpdateOrderStatus(ctx context.Context, req *UpdateOrderStatusRequest) (string, error){
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
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
		      "updatedat": updatedAt,
		    },
	  	}
		_, err = repo.orderCollection.UpdateOne(ctx, bson.M{"_id": orderId}, update)
		return req.Status, err
	}else if req.Status == "finalised" {
		update := bson.M{
		    "$set": bson.M{
		      "status": req.Status,
		      "updatedat": updatedAt,
		    },
	  	}
		_, err := repo.orderCollection.UpdateOne(ctx, bson.M{"_id": orderId}, update)
		if err != nil {
			return "", errors.New(fmt.Sprintf("update order status request failed... %v", err))
		}
		order, err := repo.GetOrder(
			ctx,
			orderId,
		)
		if err != nil {
			return "", errors.New(fmt.Sprintf("get order request failed... %v", err))
		}
		wallet, err := repo.GetWallet(
			ctx,
			&GetWalletRequest{
				UserID: order.Products[0].OwnerID,
			},
		)
		if err != nil {
			return "", errors.New(fmt.Sprintf("get wallet request failed... %v", err))
		}
		partialCancelAmount := float32(0)
		if len(order.Offers) > 0 {
			for _, offer := range order.Offers {
				if offer.Status == "accepted" {
					partialCancelAmount += offer.Amount
				}
			}
		}
		err = repo.AddTransaction(
			ctx,
			&AddTransactionRequest{
				Transaction: Transaction{
					WalletID: wallet.ID.Hex(),
					Amount: order.SubTotal + order.ShippingFees - partialCancelAmount,
					Type: "gain",
					OrderID: order.ID.Hex(),
				},
			},
		)
		if err != nil {
			return "", errors.New(fmt.Sprintf("add transaction request failed... %v", err))
		}
		return req.Status, err
	}else if req.Status == "canceled" {
		update := bson.M{
		    "$set": bson.M{
		      "status": req.Status,
		      "updatedat": updatedAt,
		    },
	  	}
		_, err := repo.orderCollection.UpdateOne(ctx, bson.M{"_id": orderId}, update)
		if err != nil {
			return "", errors.New(fmt.Sprintf("update order status request failed... %v", err))
		}
		order, err := repo.GetOrder(
			ctx,
			orderId,
		)
		if err != nil {
			return "", errors.New(fmt.Sprintf("get order request failed... %v", err))
		}
		wallet, err := repo.GetWallet(
			ctx,
			&GetWalletRequest{
				UserID: order.UserID,
			},
		)
		if err != nil {
			return "", errors.New(fmt.Sprintf("get wallet request failed... %v", err))
		}
		err = repo.AddTransaction(
			ctx,
			&AddTransactionRequest{
				Transaction: Transaction{
					WalletID: wallet.ID.Hex(),
					Amount: order.Total,
					Type: "cancel",
					OrderID: order.ID.Hex(),
				},
			},
		)
		if err != nil {
			return "", errors.New(fmt.Sprintf("add transaction request failed... %v", err))
		}
		return req.Status, err
	}else{
		update := bson.M{
		    "$set": bson.M{
		      "status": req.Status,
		      "updatedat": updatedAt,
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
	token, err := repo.GetThaiPostToken(ctx, req)
	form, _ := json.Marshal(GetThaiPostStatusRequest{
		Status: "all",
		Language: "EN",
		Barcode: []string{req.TrackID},
	})
	httpReq, err := http.NewRequest("POST", "https://trackapi.thailandpost.co.th/post/api/v1/track", bytes.NewBuffer(form))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Token "+token[1:len(token)-1])

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