// cart-service/main.go

package main

import(
	"context"
	_ "log"

	pb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	"go.mongodb.org/mongo-driver/mongo"
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

func MarshalAddToCartResponse(req *pb.AddToCartResponse) *AddToCartResponse{
	return &AddToCartResponse{
		Added: 			req.Added,
	}
}

func UnmarshalAddToCartResponse(req *AddToCartResponse) *pb.AddToCartResponse{
	return &pb.AddToCartResponse{
		Added: 			req.Added,
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

func MarshalDeleteFromCartResponse(req *pb.DeleteFromCartResponse) *DeleteFromCartResponse{
	return &DeleteFromCartResponse{
		Deleted: 			req.Deleted,
	}
}

func UnmarshalDeleteFromCartResponse(req *DeleteFromCartResponse) *pb.DeleteFromCartResponse{
	return &pb.DeleteFromCartResponse{
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

func MarshalGetResponse(req *pb.GetResponse) *GetResponse{
	return &GetResponse{
		Cart: 			*MarshalCart(req.Cart),
	}
}

func UnmarshalGetResponse(req *GetResponse) *pb.GetResponse{
	return &pb.GetResponse{
		Cart: 			UnmarshalCart(&req.Cart),
	}
}

func MarshalProduct(product *pb.Product) *Product{
	objId, _ := primitive.ObjectIDFromHex(product.Id)
	return &Product{
		ID:				objId,
		Disponible:		product.Disponible,
		Title:			product.Title,
		Price:			product.Price,
		Photos:			MarshalPhotos(product.Photos),
		Category:		*MarshalCategory(product.Category),
		Size:			product.Size,
		Color1:			*MarshalColor(product.Color1),
		Color2:			*MarshalColor(product.Color2),
		Brand:			product.Brand,
		Condition:		product.Condition,
		Material:		product.Material,
	}
}

func UnmarshalProduct(product *Product, userId string) *pb.Product{
	return &pb.Product{
		Id:				product.ID.Hex(),
		Disponible:		product.Disponible,
		Title:			product.Title,
		Price:			product.Price,
		Photos:			UnmarshalPhotos(product.Photos),
		Category:		UnmarshalCategory(&product.Category),
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

func UnmarshalCart(cart *Cart, userId string) *pb.Cart{
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
	AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartRequest, error)
	DeleteFromCart(ctx context.Context, req *pb.DeleteFromCartRequest) (*pb.DeleteFromCartRequest, error)
	GetCart(ctx context.Context, req *pb.GetRequest) (*pb.GetRequest, error)
}

type MongoRepository struct{
	collection 	*mongo.Collection
}

func (repo *MongoRepository) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartRequest, error){
	updated := append(repo.items, req)
	repo.items = updated
	return req, nil
}