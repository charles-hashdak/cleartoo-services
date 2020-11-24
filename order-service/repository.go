// orrder-service/main.go

package main

import(
	"context"
	_ "log"

	pb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID 			string 				`json:"user_id"`
	Products 		Products 			`json:"products"`
	SubTotal 		int32 				`json:"sub_total"`
	ShippingFees    int32 				`json:"shipping_fees"`
	Taxes 			int32 				`json:"taxes"`
	Total 			int32 				`json:"total"`
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

type GetResponse struct {
	Order 			Order
}

func MarshalOrderRequest(req *pb.OrderRequest) *OrderRequest{
	return &OrderRequest{
		UserID: 		req.UserId,
		Order: 			*MarshalOrder(req.Order),
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

func MarshalGetSingleRequest(req *pb.GetSingleRequest) *GetSingleRequest{
	return &GetSingleRequest{
		OrderID: 		req.OrderId,
	}
}

func UnmarshalGetSingleRequest(req *GetSingleRequest) *pb.GetSingleRequest{
	return &pb.GetSingleRequest{
		OrderId: 		req.OrderID,
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

func MarshalOrder(order *pb.Order) *Order{
	objId, _ := primitive.ObjectIDFromHex(order.Id)
	return &Order{
		ID:				objId,
		UserID:			order.UserId,
		Products:		MarshalProducts(order.Products),
	}
}

func UnmarshalOrder(order *Order) *pb.Order{
	return &pb.Order{
		Id:				order.ID.Hex(),
		UserId:			order.UserID,
		Products:		UnmarshalProducts(order.Products),
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

func UnmarshalOrderCollection(orders []*Order, userId string) []*pb.Order {
	collection := make([]*pb.Order, 0)
	for _, order := range orders {
		collection = append(collection, UnmarshalOrder(order, userId))
	}
	return collection
}

type repository interface{
	Order(ctx context.Context, req *OrderRequest) error
	GetSingleOrder(ctx context.Context, req *GetRequest) (*Order, error)
	GetOrders(ctx context.Context, req *GetRequest) ([]*Order, error)
	GetSales(ctx context.Context, req *GetRequest) ([]*Order, error)
}

type MongoRepository struct{
	orderCollection 	*mongo.Collection
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
	_, err := repo.cartCollection.InsertOne(ctx, req.Order)
	return err
}

func (repo *MongoRepository) GetOrders(ctx context.Context, req *GetRequest) ([]*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"user_id", bson.D{bson.E{"$eq", req.UserID}}})
	//bsonFilters = append(bsonFilters, bson.E{"disponible", bson.D{bson.E{"$eq", true}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	var orders []*Order
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		//order.Wishers = make([]string, 0)
		orders = append(orders, order)
	}
	return orders, err
}

func (repo *MongoRepository) GetSales(ctx context.Context, req *GetRequest) (*Order, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"products.owner_id", bson.D{bson.E{"$elemMatch", req.UserID}}})
	//bsonFilters = append(bsonFilters, bson.E{"disponible", bson.D{bson.E{"$eq", true}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.orderCollection.Find(ctx,  bsonFilters, opts)
	var orders []*Order
	for cur.Next(ctx) {
		var order *Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		//order.Wishers = make([]string, 0)
		orders = append(orders, order)
	}
	return orders, err
}