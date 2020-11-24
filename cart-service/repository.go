// cart-service/main.go

package main

import(
	"context"
	_ "log"
	"fmt"

	pb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

type Cart struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID 			string 				`json:"user_id"`
	Products 		Products 			`json:"products"`
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

type AddToCartRequest struct {
	UserID 			string
	Product 		Product
}

type AddToCartResponse struct {
	Added			bool
}

type DeleteFromCartRequest struct {
	UserID 			string
	Product 		Product
}

type DeleteFromCartResponse struct {
	Deleted			bool
}

type GetRequest struct {
	UserID 			string
}

type GetResponse struct {
	Cart 			Cart
}

type IsInCartRequest struct {
	UserID 			string
	ProductID 		string
}

type IsInCartResponse struct {
	IsInCart 		bool
}

func MarshalAddToCartRequest(req *pb.AddToCartRequest) *AddToCartRequest{
	return &AddToCartRequest{
		UserID: 		req.UserId,
		Product: 		*MarshalProduct(req.Product),
	}
}

func UnmarshalAddToCartRequest(req *AddToCartRequest) *pb.AddToCartRequest{
	return &pb.AddToCartRequest{
		UserId: 		req.UserID,
		Product: 		UnmarshalProduct(&req.Product),
	}
}

func MarshalDeleteFromCartRequest(req *pb.DeleteFromCartRequest) *DeleteFromCartRequest{
	return &DeleteFromCartRequest{
		UserID: 		req.UserId,
		Product: 		*MarshalProduct(req.Product),
	}
}

func UnmarshalDeleteFromCartRequest(req *DeleteFromCartRequest) *pb.DeleteFromCartRequest{
	return &pb.DeleteFromCartRequest{
		UserId: 		req.UserID,
		Product: 		UnmarshalProduct(&req.Product),
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

func MarshalIsInCartRequest(req *pb.IsInCartRequest) *IsInCartRequest{
	return &IsInCartRequest{
		UserID: 		req.UserId,
		ProductID: 		req.ProductId,
	}
}

func UnmarshalIsInCartRequest(req *IsInCartRequest) *pb.IsInCartRequest{
	return &pb.IsInCartRequest{
		UserId: 		req.UserID,
		ProductId: 		req.ProductID,
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

func MarshalCart(cart *pb.Cart) *Cart{
	objId, _ := primitive.ObjectIDFromHex(cart.Id)
	return &Cart{
		ID:				objId,
		UserID:			cart.UserId,
		Products:		MarshalProducts(cart.Products),
	}
}

func UnmarshalCart(cart *Cart) *pb.Cart{
	return &pb.Cart{
		Id:				cart.ID.Hex(),
		UserId:			cart.UserID,
		Products:		UnmarshalProducts(cart.Products),
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

type repository interface{
	CreateCart(ctx context.Context, req *GetRequest) error
	AddToCart(ctx context.Context, req *AddToCartRequest) error
	DeleteFromCart(ctx context.Context, req *DeleteFromCartRequest) error
	GetCart(ctx context.Context, req *GetRequest) (*Cart, error)
	IsInCart(ctx context.Context, req *IsInCartRequest) (bool, error)
}

type MongoRepository struct{
	cartCollection 	*mongo.Collection
}

func (repo *MongoRepository) AddToCart(ctx context.Context, req *AddToCartRequest) error{
	fmt.Println("req")
	fmt.Println(req)
	_, err := repo.cartCollection.UpdateOne(
	    ctx,
	    bson.M{"user_id": req.UserID},
	    bson.D{
	        {"$push", bson.D{{"products", req.Product}}},
	    },
	)
	if(err != nil){
		return err
	}
	return nil
}

func (repo *MongoRepository) DeleteFromCart(ctx context.Context, req *DeleteFromCartRequest) error{
	_, err := repo.cartCollection.UpdateOne(
	    ctx,
	    bson.M{"user_id": req.UserID},
	    bson.D{
	        {"$pull", bson.D{{"products", req.Product}}},
	    },
	)
	if(err != nil){
		return err
	}
	return nil
}

func (repo *MongoRepository) GetCart(ctx context.Context, req *GetRequest) (*Cart, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"user_id", bson.D{bson.E{"$eq", req.UserID}}})
	//bsonFilters = append(bsonFilters, bson.E{"disponible", bson.D{bson.E{"$eq", true}}})
	var cart *Cart = new(Cart)
	if err := repo.cartCollection.FindOne(ctx, bsonFilters).Decode(&cart); err != nil {
		return nil, err
	}
	return cart, nil
}

func (repo *MongoRepository) CreateCart(ctx context.Context, req *GetRequest) error{
	var cart = new(Cart)
	cart.UserID = req.UserID
	cart.Products = make(Products, 0)
	_, err := repo.cartCollection.InsertOne(ctx, cart)
	return err
}

func (repo *MongoRepository) IsInCart(ctx context.Context, req *IsInCartRequest) (bool, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"user_id", bson.D{bson.E{"$eq", req.UserID}}})
	bsonFilters = append(bsonFilters, bson.E{"products.id", bson.D{bson.E{"$matchElement", req.ProductID}}})
	if count, err := repo.cartCollection.CountDocuments(ctx, bsonFilters); err != nil {
		return nil, err
	}
	return count, nil
}