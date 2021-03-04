// orrder-service/main.go

package main

import(
	"context"
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
	SubTotal 		int32 				`json:"sub_total"`
	ShippingFees    int32 				`json:"shipping_fees"`
	Taxes 			int32 				`json:"taxes"`
	Total 			int32 				`json:"total"`
	Status 			string 				`json:"status"`
	ShippingID 		string 				`json:"shipping_id"`
	PaymentID 		string 				`json:"payment_id"`
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
	Disponible 		bool
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
	WalletID 		string
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

func MarshalProduct(product *pb.Product) *Product{
	objId, _ := primitive.ObjectIDFromHex(product.Id)
	return &Product{
		ID:				objId,
		Disponible:		product.Disponible,
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
		Disponible:		product.Disponible,
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
		ShippingID:		order.ShippingId,
		PaymentID:		order.PaymentId,
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
		ShippingId:		order.ShippingID,
		PaymentId:		order.PaymentID,
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
	//bsonFilters = append(bsonFilters, bson.E{"disponible", bson.D{bson.E{"$eq", true}}})
	var order *Order
	err := repo.orderCollection.FindOne(ctx, bsonFilters, nil).Decode(&order)
	return order, err
}

func (repo *MongoRepository) Order(ctx context.Context, req *OrderRequest) error{
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
	bsonFilters = append(bsonFilters, bson.E{"_id", bson.D{bson.E{"$eq", req.WalletID}}})
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