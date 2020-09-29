// catalog-service/repository.go

package main

import(
	"context"

	pb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	"go.mongodb.org/mongo-driver/mongo"
)

// Catalog service database structure and definitions

type Product struct{
	ID 				string 		`json:"id"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	Price 			int32 		`json:"price"`
	Photos 			Photos 		`json:"photos"`
	Category 		Category	`json:"category"`
	Size 			Size 		`json:"size"`
	Color1 			Color 		`json:"color1"`
	Color2 			Color 		`json:"color2"`
	Brand 			Brand 		`json:"brand"`
	Condition 		Condition 	`json:"condition"`
	OwnerID 		string 		`json:"owner_id"`
	WishlistCount 	int32 		`json:"wishlist_count"`
	Country 		Country 	`json:"country"`
	City 			City 		`json:"city"`
	CreatedAt 		string 		`json:"created_at"`
	UpdatedAt 		string 		`json:"updated_at"`
	ViewCount 		int32 		`json:"view_count"`
}

type Category struct{
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
	ParentCategory 	Category 	`json:"parent_category"`
	Sizes 			Sizes 		`json:"sizes"`
}

type Size struct{
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
}

type Sizes []*Size

type Photo struct{
	ID 				string 		`json:"id"`
	Url 			string 		`json:"url"`
	IsMain 			bool  		`json:"is_main"`
	Height 			int32 		`json:"height"`
	Width 			int32 		`json:"width"`
	Thumbnails 		Thumbnails 	`json:"thumbnails"`
}

type Thumbnail struct{
	ID 				string 		`json:"id"`
	Url 			string 		`json:"url"`
	Height 			int32 		`json:"height"`
	Width 			int32 		`json:"width"`
}

type Thumbnails []*Thumbnail

type Color struct{
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
	HexCode 		string 		`json:"hex_code"`
	Image 			string 		`json:"image"`
}

type Brand struct{
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
}

type Condition struct{
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
}

type Country struct{
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
}

type City struct{
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
}

// Utils functions to marshal and unmarshal between protobuff and mongodb

func MarshalProduct(product *pb.Product) *Product{
	return &Product{
		ID:				product.Id,
		Title:			product.Title,
		Description:	product.Description,
		Price:			product.Price,
		Photos:			product.Photos,
		Category:		product.Category,
		Size:			product.Size,
		Color1:			product.Color1,
		Color2:			product.Color2,
		Brand:			product.Brand,
		Condition:		product.Condition,
		OwnerID:		product.OwnerId,
		WishlistCount:	product.WishlistCount,
		Country:		product.Country,
		City:			product.City,
		CreatedAt:		product.CreatedAt,
		UpdatedAt:		product.UpdatedAt,
		ViewCount:		product.ViewCount,
	}
}

func UnmarshalProduct(product *Product) *pb.Product{
	return &pb.Product{
		Id:				product.ID,
		Title:			product.Title,
		Description:	product.Description,
		Price:			product.Price,
		Photos:			product.Photos,
		Category:		product.Category,
		Size:			product.Size,
		Color1:			product.Color1,
		Color2:			product.Color2,
		Brand:			product.Brand,
		Condition:		product.Condition,
		OwnerId:		product.OwnerID,
		WishlistCount:	product.WishlistCount,
		Country:		product.Country,
		City:			product.City,
		CreatedAt:		product.CreatedAt,
		UpdatedAt:		product.UpdatedAt,
		ViewCount:		product.ViewCount,
	}
}

func UnmarshalProductCollection(products []*Product) []*pb.Product {
	collection := make([]*pb.Product, 0)
	for _, product := range products {
		collection = append(collection, UnmarshalProduct(product))
	}
	return collection
}



type repository interface{
	CreateProduct(ctx context.Context, product *pb.Product) error
	GetProducts(ctx context.Context, req *pb.GetProductsRequest) ([]*pb.Product, error)
}

type MongoRepository struct{
	collection *mongo.Collection
}

func (repo *Repository) CreateProduct(ctx context.Context, product *pb.Product) error{
	_, err := repository.collection.InsertOne(ctx, product)
	return err
}

func (repo *Repository) GetProducts(ctx context.Context, req *pb.GetProductsRequest) ([]*pb.Product, error){
	filters := req.GetFilters()
	bsonFilters := bson.M{}
	for _, f := range filters {
		bsonFilters[f.Key] = f.Value
	}
	cur, err := repository.collection.Find(ctx, bsonFilters, nil)
	var products []*Product
	for cur.Next(ctx) {
		var product *Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}