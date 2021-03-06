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
	OwnerID 		string
	Weight 			int32
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
	ProductID 		string
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

type CleanCartsFromProductRequest struct {
	ProductID 		string
}

type CleanCartsFromProductResponse struct {
	Cleaned 		bool
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
		ProductID: 		req.ProductId,
	}
}

func UnmarshalDeleteFromCartRequest(req *DeleteFromCartRequest) *pb.DeleteFromCartRequest{
	return &pb.DeleteFromCartRequest{
		UserId: 		req.UserID,
		ProductId: 		req.ProductID,
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

func MarshalCleanCartsFromProductRequest(req *pb.CleanCartsFromProductRequest) *CleanCartsFromProductRequest{
	return &CleanCartsFromProductRequest{
		ProductID: 		req.ProductId,
	}
}

func UnmarshalCleanCartsFromProductRequest(req *CleanCartsFromProductRequest) *pb.CleanCartsFromProductRequest{
	return &pb.CleanCartsFromProductRequest{
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
		OwnerID:		product.OwnerId,
		Weight:			product.Weight,
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
		Weight:			product.Weight,
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
	CleanCartsFromProduct(ctx context.Context, req *CleanCartsFromProductRequest) error
	EmptyCart(ctx context.Context, req *GetRequest) error
}

type MongoRepository struct{
	cartCollection 	*mongo.Collection
}

func (repo *MongoRepository) AddToCart(ctx context.Context, req *AddToCartRequest) error{
	bsonFilters := bson.D{
		{"userid", req.UserID},
	}
	count, countErr := repo.cartCollection.CountDocuments(ctx, bsonFilters)
	if countErr != nil {
		return countErr
	}
	if count == 0 {
		createErr := repo.CreateCart(ctx, &GetRequest{UserID: req.UserID})
		if createErr != nil {
			return createErr
		}
	}
	// cart, cartErr := repo.GetCart(ctx, &GetRequest{UserID: req.UserID})
	// if cartErr != nil {
	// 	return cartErr
	// }
	// var valid = true
	// for _, product := range cart.Products {
	// 	if product.OwnerID != req.UserID {
	// 		valid = false
	// 	}
	// }
	// if !valid {
	// 	emptyErr := repo.EmptyCart(ctx, &GetRequest{UserID: req.UserID})
	// 	if emptyErr != nil {
	// 		return emptyErr
	// 	}
	// }
	_, err := repo.cartCollection.UpdateOne(
	    ctx,
	    bson.M{"userid": req.UserID},
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
	productId, _ := primitive.ObjectIDFromHex(req.ProductID)
	_, err := repo.cartCollection.UpdateOne(
	    ctx,
	    bson.M{"userid": req.UserID},
	    bson.D{
	        {"$pull", bson.D{
	        	{"products", bson.D{
	        		{"id", productId},
	        	}},
	        }},
	    },
	)
	if(err != nil){
		return err
	}
	return nil
}

func (repo *MongoRepository) GetCart(ctx context.Context, req *GetRequest) (*Cart, error){
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"userid", bson.D{bson.E{"$eq", req.UserID}}})
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
	productId, _ := primitive.ObjectIDFromHex(req.ProductID)
	bsonFilters := bson.D{
		{"userid", req.UserID},
		{"products", bson.D{
			{"$elemMatch", bson.D{
				{"id", productId},
			}},
		}},
	}
	count, err := repo.cartCollection.CountDocuments(ctx, bsonFilters)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (repo *MongoRepository) CleanCartsFromProduct(ctx context.Context, req *CleanCartsFromProductRequest) error{
	productId, _ := primitive.ObjectIDFromHex(req.ProductID)
	bsonFilters := bson.D{
		{"products", bson.D{
			{"$elemMatch", bson.D{
				{"id", productId},
			}},
		}},
	}
	_, err := repo.cartCollection.UpdateMany(
	    ctx,
	    bsonFilters,
	    bson.D{
	        {"$pull", bson.D{
	        	{"products", bson.D{
	        		{"id", productId},
	        	}},
	        }},
	    },
	)
	if(err != nil){
		return err
	}
	return nil
}

func (repo *MongoRepository) EmptyCart(ctx context.Context, req *GetRequest) error{
	fmt.Println(req.UserID)
	cartProducts := make(Products, 0)
	_, err := repo.cartCollection.UpdateOne(
	    ctx,
	    bson.D{
			{"userid", req.UserID},
		},
	    bson.D{
	        {"$set", bson.D{
	        	{"products", cartProducts},
	        }},
	    },
	)
	if(err != nil){
		return err
	}
	return nil
}