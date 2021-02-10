// catalog-service/repository.go

package main

import(
	"context"
	"time"
	"strings"

	pb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Catalog service database structure and definitions

type Product struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Available 		bool 				`json:"available"`
	Title 			string 				`json:"title"`
	Description 	string 				`json:"description"`
	Price 			int32 				`json:"price"`
	Photos 			Photos 				`json:"photos"`
	Gender  		Gender 				`json:"gender"`
	Category 		Category			`json:"category"`
	Size 			string 				`json:"size"`
	Color1 			Color 				`json:"color1"`
	Color2 			Color 				`json:"color2"`
	Brand 			string 				`json:"brand"`
	Condition 		string 				`json:"condition"`
	Material 		string 				`json:"material"`
	Owner 			Owner 				`json:"owner"`
	Wishers 		[]string 			`json:"Wishers"`
	WishlistCount 	int32 				`json:"wishlist_count"`
	Country 		string 				`json:"country"`
	City 			string 				`json:"city"`
	CreatedAt 		string 				`json:"created_at"`
	UpdatedAt 		string 				`json:"updated_at"`
	ViewCount 		int32 				`json:"view_count"`
	Offers 			Offers 				`json:"offers"`
	InCart 			bool
}

type Gender struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
}

type Category struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
	ParentID 		primitive.ObjectID	`bson:"parent_id,omitempty"`
	Sizes 			Sizes 				`json:"sizes"`
	Genders 		[]string 			`json:"genders"`
}

type Owner struct{
	OwnerID 		string 				`json:"owner_id"`
	Username 		string 				`json:"username"`
	Rating 			string 				`json:"rating"`
	Avatar 			Photo 				`json:"avatar"`
}

type Size struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
}

type Sizes []*Size

type Photo struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Url 			string 				`json:"url"`
	IsMain 			bool  				`json:"is_main"`
	Height 			int32 				`json:"height"`
	Width 			int32 				`json:"width"`
	Thumbnails 		Thumbnails 			`json:"thumbnails"`
}

type Photos []*Photo

type Offer struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID 			string 				`json:"user_id"`
	Amount			int32 				`json:"amount"`
	Status 			string 				`json:"status"`
}

type Offers []*Offer

type GetRequest struct{
	Filters 		Filters
	UserID 			string
	ProductID 		string
	Wished 			bool
	Limit			int64
	Offset			int64
}

type Request struct{}

type Filter struct{
	Key 			string
	Value 			string
	Condition		string
	Hex 			bool
}

type Filters []*Filter

type Thumbnail struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Url 			string 				`json:"url"`
	Height 			int32 				`json:"height"`
	Width 			int32 				`json:"width"`
}

type Thumbnails []*Thumbnail

type Color struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
	HexCode 		string 				`json:"hex_code"`
	Image 			string 				`json:"image"`
}

type Brand struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
}

type Condition struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
}

type Material struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
}

type Country struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
}

type City struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	Name 			string 				`json:"name"`
	CountryID 		string 				`json:"country_id"`
}

// Utils functions to marshal and unmarshal between protobuff and mongodb

func MarshalProduct(product *pb.Product) *Product{
	objId, _ := primitive.ObjectIDFromHex(product.Id)
	return &Product{
		ID:				objId,
		Available:		product.Available,
		Title:			product.Title,
		Description:	product.Description,
		Price:			product.Price,
		Photos:			MarshalPhotos(product.Photos),
		Gender:			*MarshalGender(product.Gender),
		Category:		*MarshalCategory(product.Category),
		Size:			product.Size,
		Color1:			*MarshalColor(product.Color1),
		Color2:			*MarshalColor(product.Color2),
		Brand:			product.Brand,
		Condition:		product.Condition,
		Material:		product.Material,
		Owner:			*MarshalOwner(product.Owner),
		Wishers:		product.Wishers,
		WishlistCount:	product.WishlistCount,
		Country:		product.Country,
		City:			product.City,
		CreatedAt:		product.CreatedAt,
		UpdatedAt:		product.UpdatedAt,
		ViewCount:		product.ViewCount,
		Offers:			MarshalOffers(product.Offers),
	}
}

func UnmarshalProduct(product *Product, userId string) *pb.Product{
	var wished = false;
	for _, b := range product.Wishers {
        if b == userId {
            wished = true
        }
    }
	return &pb.Product{
		Id:				product.ID.Hex(),
		Available:		product.Available,
		Title:			product.Title,
		Description:	product.Description,
		Price:			product.Price,
		Photos:			UnmarshalPhotos(product.Photos),
		Gender:			UnmarshalGender(&product.Gender),
		Category:		UnmarshalCategory(&product.Category),
		Size:			product.Size,
		Color1:			UnmarshalColor(&product.Color1),
		Color2:			UnmarshalColor(&product.Color2),
		Brand:			product.Brand,
		Condition:		product.Condition,
		Material:		product.Material,
		Owner:			UnmarshalOwner(&product.Owner),
		WishlistCount:	product.WishlistCount,
		Country:		product.Country,
		City:			product.City,
		CreatedAt:		product.CreatedAt,
		UpdatedAt:		product.UpdatedAt,
		ViewCount:		product.ViewCount,
		Wished:			wished,
		Offers:			UnmarshalOffers(product.Offers),
		InCart:			product.InCart,
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
		Thumbnails:		MarshalThumbnails(photo.Thumbnails),
	}
}

func MarshalPhotos(photos []*pb.Photo) Photos {
	collection := make(Photos, 0)
	for _, photo := range photos {
		collection = append(collection, MarshalPhoto(photo))
	}
	return collection
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
		Thumbnails:		UnmarshalThumbnails(photo.Thumbnails),
	}
}

func UnmarshalPhotos(photos Photos) []*pb.Photo {
	collection := make([]*pb.Photo, 0)
	for _, photo := range photos {
		collection = append(collection, UnmarshalPhoto(photo))
	}
	return collection
}

func MarshalOffer(offer *pb.Offer) *Offer{
	if(offer == nil){
		return &Offer{}
	}
	objId, _ := primitive.ObjectIDFromHex(offer.Id)
	return &Offer{
		ID:				objId,
		UserID:			offer.UserId,
		Amount:			offer.Amount,
		Status:			offer.Status,
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
		UserId:			offer.UserID,
		Amount:			offer.Amount,
		Status:			offer.Status,
	}
}

func UnmarshalOffers(offers Offers) []*pb.Offer {
	collection := make([]*pb.Offer, 0)
	for _, offer := range offers {
		collection = append(collection, UnmarshalOffer(offer))
	}
	return collection
}

func MarshalThumbnail(thumbnail *pb.Thumbnail) *Thumbnail{
	if(thumbnail == nil){
		return &Thumbnail{}
	}
	objId, _ := primitive.ObjectIDFromHex(thumbnail.Id)
	return &Thumbnail{
		ID:				objId,
		Url: 			thumbnail.Url,
		Height:			thumbnail.Height,
		Width:			thumbnail.Width,
	}
}

func MarshalThumbnails(thumbnails []*pb.Thumbnail) Thumbnails {
	collection := make(Thumbnails, 0)
	for _, thumbnail := range thumbnails {
		collection = append(collection, MarshalThumbnail(thumbnail))
	}
	return collection
}

func UnmarshalThumbnail(thumbnail *Thumbnail) *pb.Thumbnail{
	if(thumbnail == nil){
		return &pb.Thumbnail{}
	}
	return &pb.Thumbnail{
		Id:				thumbnail.ID.Hex(),
		Url: 			thumbnail.Url,
		Height:			thumbnail.Height,
		Width:			thumbnail.Width,
	}
}

func UnmarshalThumbnails(thumbnails Thumbnails) []*pb.Thumbnail {
	collection := make([]*pb.Thumbnail, 0)
	for _, thumbnail := range thumbnails {
		collection = append(collection, UnmarshalThumbnail(thumbnail))
	}
	return collection
}

func MarshalCategory(category *pb.Category) *Category{
	if(category == nil){
		return &Category{}
	}
	objId, _ := primitive.ObjectIDFromHex(category.Id)
	parentObjId, _ := primitive.ObjectIDFromHex(category.ParentId)
	return &Category{
		ID:				objId,
		Name:			category.Name,
		ParentID:		parentObjId,
		Sizes:			MarshalSizes(category.Sizes),
		Genders:		category.Genders,
	}
}

func UnmarshalCategory(category *Category) *pb.Category{
	if(category == nil){
		return &pb.Category{}
	}
	return &pb.Category{
		Id:				category.ID.Hex(),
		Name:			category.Name,
		ParentId:		category.ParentID.Hex(),
		Sizes:			UnmarshalSizes(category.Sizes),
		Genders:		category.Genders,
	}
}

func MarshalGender(gender *pb.Gender) *Gender{
	if(gender == nil){
		return &Gender{}
	}
	objId, _ := primitive.ObjectIDFromHex(gender.Id)
	return &Gender{
		ID:				objId,
		Name:			gender.Name,
	}
}

func UnmarshalGender(gender *Gender) *pb.Gender{
	if(gender == nil){
		return &pb.Gender{}
	}
	return &pb.Gender{
		Id:				gender.ID.Hex(),
		Name:			gender.Name,
	}
}

func MarshalOwner(owner *pb.Owner) *Owner{
	if(owner == nil){
		return &Owner{}
	}
	return &Owner{
		OwnerID:		owner.OwnerId,
		Username:		owner.Username,
		Rating:			owner.Rating,
		Avatar:			*MarshalPhoto(owner.Avatar),
	}
}

func UnmarshalOwner(owner *Owner) *pb.Owner{
	if(owner == nil){
		return &pb.Owner{}
	}
	return &pb.Owner{
		OwnerId:		owner.OwnerID,
		Username:		owner.Username,
		Rating:			owner.Rating,
		Avatar:			UnmarshalPhoto(&owner.Avatar),
	}
}

func MarshalSize(size *pb.Size) *Size{
	if(size == nil){
		return &Size{}
	}
	objId, _ := primitive.ObjectIDFromHex(size.Id)
	return &Size{
		ID:				objId,
		Name:			size.Name,
	}
}

func MarshalSizes(sizes []*pb.Size) Sizes {
	collection := make(Sizes, 0)
	for _, size := range sizes {
		collection = append(collection, MarshalSize(size))
	}
	return collection
}

func UnmarshalSize(size *Size) *pb.Size{
	if(size == nil){
		return &pb.Size{}
	}
	return &pb.Size{
		Id:				size.ID.Hex(),
		Name:			size.Name,
	}
}

func UnmarshalSizes(sizes Sizes) []*pb.Size {
	collection := make([]*pb.Size, 0)
	for _, size := range sizes {
		collection = append(collection, UnmarshalSize(size))
	}
	return collection
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

func MarshalBrand(brand *pb.Brand) *Brand{
	if(brand == nil){
		return &Brand{}
	}
	objId, _ := primitive.ObjectIDFromHex(brand.Id)
	return &Brand{
		ID:				objId,
		Name:			brand.Name,
	}
}

func UnmarshalBrand(brand *Brand) *pb.Brand{
	if(brand == nil){
		return &pb.Brand{}
	}
	return &pb.Brand{
		Id:				brand.ID.Hex(),
		Name:			brand.Name,
	}
}

func MarshalCondition(condition *pb.Condition) *Condition{
	if(condition == nil){
		return &Condition{}
	}
	objId, _ := primitive.ObjectIDFromHex(condition.Id)
	return &Condition{
		ID:				objId,
		Name:			condition.Name,
	}
}

func UnmarshalCondition(condition *Condition) *pb.Condition{
	if(condition == nil){
		return &pb.Condition{}
	}
	return &pb.Condition{
		Id:				condition.ID.Hex(),
		Name:			condition.Name,
	}
}

func MarshalMaterial(material *pb.Material) *Material{
	if(material == nil){
		return &Material{}
	}
	objId, _ := primitive.ObjectIDFromHex(material.Id)
	return &Material{
		ID:				objId,
		Name:			material.Name,
	}
}

func UnmarshalMaterial(material *Material) *pb.Material{
	if(material == nil){
		return &pb.Material{}
	}
	return &pb.Material{
		Id:				material.ID.Hex(),
		Name:			material.Name,
	}
}

func MarshalCountry(country *pb.Country) *Country{
	if(country == nil){
		return &Country{}
	}
	objId, _ := primitive.ObjectIDFromHex(country.Id)
	return &Country{
		ID:				objId,
		Name:			country.Name,
	}
}

func UnmarshalCountry(country *Country) *pb.Country{
	if(country == nil){
		return &pb.Country{}
	}
	return &pb.Country{
		Id:				country.ID.Hex(),
		Name:			country.Name,
	}
}

func MarshalCity(city *pb.City) *City{
	if(city == nil){
		return &City{}
	}
	objId, _ := primitive.ObjectIDFromHex(city.Id)
	return &City{
		ID:				objId,
		Name:			city.Name,
		CountryID:		city.CountryId,
	}
}

func UnmarshalCity(city *City) *pb.City{
	if(city == nil){
		return &pb.City{}
	}
	return &pb.City{
		Id:				city.ID.Hex(),
		Name:			city.Name,
		CountryId:		city.CountryID,
	}
}

func MarshalGetRequest(req *pb.GetRequest) *GetRequest{
	return &GetRequest{
		Filters:		MarshalFilters(req.Filters),
		UserID: 		req.UserId,
		ProductID: 		req.ProductId,
		Wished: 		req.Wished,
		Limit: 			req.Limit,
		Offset:			req.Offset,
	}
}

func UnmarshalGetRequest(req *GetRequest) *pb.GetRequest{
	return &pb.GetRequest{
		Filters:		UnmarshalFilters(req.Filters),
		UserId: 		req.UserID,
		ProductId: 		req.ProductID,
		Wished: 		req.Wished,
		Limit: 			req.Limit,
		Offset:			req.Offset,
	}
}

func MarshalRequest(req *pb.Request) *Request{
	return &Request{}
}

func UnmarshalRequest(req *Request) *pb.Request{
	return &pb.Request{}
}

func UnmarshalProductCollection(products []*Product, userId string) []*pb.Product {
	collection := make([]*pb.Product, 0)
	for _, product := range products {
		collection = append(collection, UnmarshalProduct(product, userId))
	}
	return collection
}

func UnmarshalSizeCollection(sizes []*Size) []*pb.Size {
	collection := make([]*pb.Size, 0)
	for _, size := range sizes {
		collection = append(collection, UnmarshalSize(size))
	}
	return collection
}

func UnmarshalGenderCollection(genders []*Gender) []*pb.Gender {
	collection := make([]*pb.Gender, 0)
	for _, gender := range genders {
		collection = append(collection, UnmarshalGender(gender))
	}
	return collection
}

func UnmarshalCategoryCollection(categories []*Category) []*pb.Category {
	collection := make([]*pb.Category, 0)
	for _, category := range categories {
		collection = append(collection, UnmarshalCategory(category))
	}
	return collection
}

func UnmarshalBrandCollection(brands []*Brand) []*pb.Brand {
	collection := make([]*pb.Brand, 0)
	for _, brand := range brands {
		collection = append(collection, UnmarshalBrand(brand))
	}
	return collection
}

func UnmarshalColorCollection(colors []*Color) []*pb.Color {
	collection := make([]*pb.Color, 0)
	for _, color := range colors {
		collection = append(collection, UnmarshalColor(color))
	}
	return collection
}

func UnmarshalConditionCollection(conditions []*Condition) []*pb.Condition {
	collection := make([]*pb.Condition, 0)
	for _, condition := range conditions {
		collection = append(collection, UnmarshalCondition(condition))
	}
	return collection
}

func UnmarshalMaterialCollection(materials []*Material) []*pb.Material {
	collection := make([]*pb.Material, 0)
	for _, material := range materials {
		collection = append(collection, UnmarshalMaterial(material))
	}
	return collection
}

func MarshalFilter(filter *pb.Filter) *Filter{
	return &Filter{
		Key:				filter.Key,
		Value:				filter.Value,
		Condition:			filter.Condition,
		Hex:				filter.Hex,
	}
}

func MarshalFilters(filters []*pb.Filter) Filters {
	collection := make(Filters, 0)
	for _, filter := range filters {
		collection = append(collection, MarshalFilter(filter))
	}
	return collection
}

func UnmarshalFilter(filter *Filter) *pb.Filter{
	return &pb.Filter{
		Key:				filter.Key,
		Value:				filter.Value,
		Condition:			filter.Condition,
		Hex:				filter.Hex,
	}
}

func UnmarshalFilters(filters Filters) []*pb.Filter {
	collection := make([]*pb.Filter, 0)
	for _, filter := range filters {
		collection = append(collection, UnmarshalFilter(filter))
	}
	return collection
}



type repository interface{
	CreateProduct(ctx context.Context, product *Product) error
	EditProduct(ctx context.Context, product *Product) error
	GetProducts(ctx context.Context, req *GetRequest) ([]*Product, error)
	GetProduct(ctx context.Context, req *GetRequest) (*Product, error)
	Wish(ctx context.Context, req *GetRequest) error
	GetWishes(ctx context.Context, req *GetRequest) ([]*Product, error)
	GetSizes(ctx context.Context, req *GetRequest) ([]*Size, error)
	GetGenders(ctx context.Context, req *Request) ([]*Gender, error)
	GetCategories(ctx context.Context, req *GetRequest) ([]*Category, error)
	GetBrands(ctx context.Context, req *Request) ([]*Brand, error)
	GetColors(ctx context.Context, req *Request) ([]*Color, error)
	GetConditions(ctx context.Context, req *Request) ([]*Condition, error)
	GetMaterials(ctx context.Context, req *Request) ([]*Material, error)
}

type MongoRepository struct{
	productsCollection *mongo.Collection
	gendersCollection *mongo.Collection
	categoriesCollection *mongo.Collection
	sizesCollection *mongo.Collection
	brandsCollection *mongo.Collection
	colorsCollection *mongo.Collection
	conditionsCollection *mongo.Collection
	materialsCollection *mongo.Collection
}

func (repo *MongoRepository) CreateProduct(ctx context.Context, product *Product) error{
	product.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	product.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	product.Wishers = []string{}
	product.WishlistCount = 0
	product.Offers = Offers{}
	_, err := repo.productsCollection.InsertOne(ctx, product)
	return err
}

func (repo *MongoRepository) EditProduct(ctx context.Context, product *Product) error{
	product.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := repo.productsCollection.UpdateOne(ctx,bson.M{"_id": product.ID} , product)
	return err
}

func (repo *MongoRepository) GetProducts(ctx context.Context, req *GetRequest) ([]*Product, error){
	filters := req.Filters
	bsonFilters := bson.D{}
	for _, f := range filters {
		if(f.Condition == ""){
			f.Condition = "$eq";
		}
		if(f.Condition == "$in"){
			newValue := strings.Split(f.Value, ",")
			if(f.Hex == true){
				var finalValue []primitive.ObjectID
				for _, v := range newValue {
					objId, _ := primitive.ObjectIDFromHex(v)
					finalValue = append(finalValue, objId)
				}
				bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, finalValue}}})
			}else{
				finalValue := newValue
				bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, finalValue}}})
			}
		}else{
			if(f.Hex == true){
				objId, _ := primitive.ObjectIDFromHex(f.Value)
				bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, objId}}})
			}else if(f.Value == "true"){
				bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, true}}})
			}else if(f.Value == "false"){
				bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, false}}})
			}else{
				bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, f.Value}}})
			}
		}
	}
	//bsonFilters = append(bsonFilters, bson.E{"available", bson.D{bson.E{"$eq", true}}})
	cur, err := repo.productsCollection.Find(ctx,  bsonFilters, options.Find().SetShowRecordID(true), options.Find().SetLimit(req.Limit), options.Find().SetSkip(req.Offset))
	if err != nil {
		return nil, err
	}
	var products []*Product
	for cur.Next(ctx) {
		var product *Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		//product.Wishers = make([]string, 0)
		products = append(products, product)
	}
	return products, err
}

func (repo *MongoRepository) GetProduct(ctx context.Context, req *GetRequest) (*Product, error){
	productId, _ := primitive.ObjectIDFromHex(req.ProductID)
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"_id", bson.D{bson.E{"$eq", productId}}})
	//bsonFilters = append(bsonFilters, bson.E{"available", bson.D{bson.E{"$eq", true}}})
	var product *Product
	err := repo.productsCollection.FindOne(ctx, bsonFilters, nil).Decode(&product)
	return product, err
}

func (repo *MongoRepository) GetWishes(ctx context.Context, req *GetRequest) ([]*Product, error){
	userId := req.UserID
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"wishers", bson.D{bson.E{"$elemMatch", bson.D{bson.E{"$eq", userId}}}}})
	//bsonFilters = append(bsonFilters, bson.E{"Available", bson.D{bson.E{"$eq", true}}})
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.productsCollection.Find(ctx, bsonFilters, opts)
	var products []*Product
	for cur.Next(ctx) {
		var product *Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		//product.Wishers = make([]string, 0)
		products = append(products, product)
	}
	return products, err
}

func (repo *MongoRepository) Wish(ctx context.Context, req *GetRequest) error{
	productId, _ := primitive.ObjectIDFromHex(req.ProductID)
	var op string
	if(req.Wished){
		op = "$push"
	}else{
		op = "$pull"
	}
	_, err := repo.productsCollection.UpdateOne(
	    ctx,
	    bson.M{"_id": productId},
	    bson.D{
	        {op, bson.D{{"wishers", req.UserID}}},
	    },
	)
	if(err != nil){
		return err
	}
	return nil
}

func (repo *MongoRepository) GetSizes(ctx context.Context, req *GetRequest) ([]*Size, error){
	filters := req.Filters
	bsonFilters := bson.M{}
	for _, f := range filters {
		bsonFilters[f.Key] = f.Value
	}
	cur, err := repo.sizesCollection.Find(ctx, bsonFilters, nil)
	var sizes []*Size
	for cur.Next(ctx) {
		var size *Size
		if err := cur.Decode(&size); err != nil {
			return nil, err
		}
		sizes = append(sizes, size)
	}
	return sizes, err
}

func (repo *MongoRepository) CreateCategory(ctx context.Context, category *pb.Category) error{
	_, err := repo.categoriesCollection.InsertOne(ctx, category)
	return err
}

func (repo *MongoRepository) GetGenders(ctx context.Context, req *Request) ([]*Gender, error){
	bsonFilters := bson.M{}
	cur, err := repo.gendersCollection.Find(ctx, bsonFilters, nil)
	var genders []*Gender
	for cur.Next(ctx) {
		var gender *Gender
		if err := cur.Decode(&gender); err != nil {
			return nil, err
		}
		genders = append(genders, gender)
	}
	return genders, err
}

func (repo *MongoRepository) GetCategories(ctx context.Context, req *GetRequest) ([]*Category, error){
	filters := req.Filters
	bsonFilters := bson.D{}
	for _, f := range filters {
		if(f.Condition == ""){
			f.Condition = "$eq";
		}
		if(f.Hex == true){
			objId, _ := primitive.ObjectIDFromHex(f.Value)
			bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, objId}}})
		}else if(f.Value == "true"){
			bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, true}}})
		}else if(f.Value == "false"){
			bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, false}}})
		}else{
			bsonFilters = append(bsonFilters, bson.E{f.Key, bson.D{bson.E{f.Condition, f.Value}}})
		}
	}
	opts := options.Find().SetShowRecordID(true)
	cur, err := repo.categoriesCollection.Find(ctx, bsonFilters, opts)
	var categories []*Category
	for cur.Next(ctx) {
		var category *Category
		if err := cur.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, err
}

func (repo *MongoRepository) GetBrands(ctx context.Context, req *Request) ([]*Brand, error){
	bsonFilters := bson.M{}
	cur, err := repo.brandsCollection.Find(ctx, bsonFilters, nil)
	var brands []*Brand
	for cur.Next(ctx) {
		var brand *Brand
		if err := cur.Decode(&brand); err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}
	return brands, err
}

func (repo *MongoRepository) GetColors(ctx context.Context, req *Request) ([]*Color, error){
	bsonFilters := bson.M{}
	cur, err := repo.colorsCollection.Find(ctx, bsonFilters, nil)
	var colors []*Color
	for cur.Next(ctx) {
		var color *Color
		if err := cur.Decode(&color); err != nil {
			return nil, err
		}
		colors = append(colors, color)
	}
	return colors, err
}

func (repo *MongoRepository) GetConditions(ctx context.Context, req *Request) ([]*Condition, error){
	bsonFilters := bson.M{}
	cur, err := repo.conditionsCollection.Find(ctx, bsonFilters, nil)
	var conditions []*Condition
	for cur.Next(ctx) {
		var condition *Condition
		if err := cur.Decode(&condition); err != nil {
			return nil, err
		}
		conditions = append(conditions, condition)
	}
	return conditions, err
}

func (repo *MongoRepository) GetMaterials(ctx context.Context, req *Request) ([]*Material, error){
	bsonFilters := bson.M{}
	cur, err := repo.materialsCollection.Find(ctx, bsonFilters, nil)
	var materials []*Material
	for cur.Next(ctx) {
		var material *Material
		if err := cur.Decode(&material); err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, err
}