// catalog-service/repository.go

package main

import(
	"context"
	"time"

	pb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Catalog service database structure and definitions

type Address struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID 			string				`json:"user_id"`
	Title 			string 				`json:"title"`
	Indications 	string 				`json:"indications"`
	IsMain 			bool 				`json:"is_main"`
	AddressLine1 	string 				`json:"address_line1"`
	AddressLine2 	string 				`json:"address_line2"`
	Country 		Country 			`json:"country"`
	City 			City 				`json:"city"`
	PostalCode		string 				`json:"postal_code"`
	CreatedAt 		string 				`json:"created_at"`
	UpdatedAt 		string 				`json:"updated_at"`
}

type Shipment struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	OrderID 		string				`json:"order_id"`
	Status 			string 				`json:"status"`
	AddressID 		string				`json:"address_id"`
	Method 			Method 				`json:"method"`
	TrackUrl 		string 				`json:"track_url"`
}

type Method struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	OrderID 		string				`json:"order_id"`
	Countries 		Countries 			`json:"countries"`
	Cities 			Cities 				`json:"cities"`
}

type Countries []*Country

type Cities []*City

type GetRequest struct{
	Filters 		Filters
	UserID 			string
	ProductID 		string
	Wished 			bool
}

type Request struct{}

type Filter struct{
	Key 			string
	Value 			string
	Condition		string
	Hex 			bool
}

type Filters []*Filter

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

func MarshalAddress(address *pb.Address) *Address{
	objId, _ := primitive.ObjectIDFromHex(address.Id)
	return &Address{
		ID: 			objId,
		UserID: 		address.UserId,
		Title: 			address.Title,
		Indications: 	address.Indications,
		IsMain: 		address.IsMain,
		AddressLine1: 	address.AddressLine1,
		AddressLine2: 	address.AddressLine2,
		Country: 		*MarshalCountry(address.Country),
		City: 			*MarshalCity(address.City),
		PostalCode:		address.PostalCode,
		CreatedAt: 		address.CreatedAt,
		UpdatedAt: 		address.UpdatedAt,
	}
}

func UnmarshalAddress(address *Address) *pb.Address{
	return &pb.Address{
		Id: 			address.ID.Hex(),
		UserId: 		address.UserID,
		Title: 			address.Title,
		Indications: 	address.Indications,
		IsMain: 		address.IsMain,
		AddressLine1: 	address.AddressLine1,
		AddressLine2: 	address.AddressLine2,
		Country: 		UnmarshalCountry(&address.Country),
		City: 			UnmarshalCity(&address.City),
		PostalCode:		address.PostalCode,
		CreatedAt: 		address.CreatedAt,
		UpdatedAt: 		address.UpdatedAt,
	}
}

func MarshalShipment(shipment *pb.Shipment) *Shipment{
	objId, _ := primitive.ObjectIDFromHex(shipment.Id)
	return &Shipment{
		ID: 			objId,
		OrderID: 		shipment.OrderId,
		Status: 		shipment.Status,
		AddressID: 		shipment.AddressId,
		Method: 		*MarshalMethod(shipment.Method),
		TrackUrl: 		shipment.TrackUrl,
	}
}

func UnmarshalShipment(shipment *Shipment) *pb.Shipment{
	return &pb.Shipment{
		Id: 			shipment.ID.Hex(),
		OrderId: 		shipment.OrderID,
		Status: 		shipment.Status,
		AddressId: 		shipment.AddressID,
		Method: 		UnmarshalMethod(&shipment.Method),
		TrackUrl: 		shipment.TrackUrl,
	}
}

func MarshalMethod(method *pb.Method) *Method{
	objId, _ := primitive.ObjectIDFromHex(method.Id)
	return &Method{
		ID: 			objId,
		OrderID: 		method.OrderId,
		Countries: 		MarshalCountries(method.Countries),
		Cities: 		MarshalCities(method.Cities),
	}
}

func UnmarshalMethod(method *Method) *pb.Method{
	return &pb.Method{
		Id: 			method.ID.Hex(),
		OrderId: 		method.OrderID,
		Countries: 		UnmarshalCountries(method.Countries),
		Cities: 		UnmarshalCities(method.Cities),
	}
}

func MarshalCountries(countries []*pb.Country) Countries {
	collection := make(Countries, 0)
	for _, country := range countries {
		collection = append(collection, MarshalCountry(country))
	}
	return collection
}

func UnmarshalCountries(countries Countries) []*pb.Country {
	collection := make([]*pb.Country, 0)
	for _, country := range countries {
		collection = append(collection, UnmarshalCountry(country))
	}
	return collection
}

func MarshalCities(cities []*pb.City) Cities {
	collection := make(Cities, 0)
	for _, city := range cities {
		collection = append(collection, MarshalCity(city))
	}
	return collection
}

func UnmarshalCities(cities Cities) []*pb.City {
	collection := make([]*pb.City, 0)
	for _, city := range cities {
		collection = append(collection, UnmarshalCity(city))
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
		AddressID: 		req.AddressId,
	}
}

func UnmarshalGetRequest(req *GetRequest) *pb.GetRequest{
	return &pb.GetRequest{
		Filters:		UnmarshalFilters(req.Filters),
		UserId: 		req.UserID,
		AddressId: 		req.AddressID,
	}
}

func MarshalRequest(req *pb.Request) *Request{
	return &Request{}
}

func UnmarshalRequest(req *Request) *pb.Request{
	return &pb.Request{}
}

func UnmarshalAddressCollection(addresses []*Address, userId string) []*pb.Address {
	collection := make([]*pb.Address, 0)
	for _, address := range addresses {
		collection = append(collection, UnmarshalAddress(address, userId))
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
	CreateAddress(ctx context.Context, address *Address) error
	GetAddresses(ctx context.Context, req *GetRequest) ([]*Address, error)
	GetAddress(ctx context.Context, req *GetRequest) (*Address, error)
	GetCountries(ctx context.Context, req *Request) ([]*Country, error)
	GetCities(ctx context.Context, req *GetRequest) ([]*City, error)
}

type MongoRepository struct{
	addressesCollection *mongo.Collection
	shipmentsCollection *mongo.Collection
	methodsCollection *mongo.Collection
	countriesCollection *mongo.Collection
	citiesCollection *mongo.Collection
}

func (repo *MongoRepository) CreateAddress(ctx context.Context, address *Address) error{
	address.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	address.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := repo.addressesCollection.InsertOne(ctx, address)
	return err
}

func (repo *MongoRepository) GetAddresses(ctx context.Context, req *GetRequest) ([]*Address, error){
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
	cur, err := repo.addressesCollection.Find(ctx,  bsonFilters, opts)
	var addresses []*Address
	for cur.Next(ctx) {
		var address *Address
		if err := cur.Decode(&address); err != nil {
			return nil, err
		}
		//address.Wishers = make([]string, 0)
		addresses = append(addresses, address)
	}
	return addresses, err
}

func (repo *MongoRepository) GetAddress(ctx context.Context, req *GetRequest) (*Address, error){
	addressId, _ := primitive.ObjectIDFromHex(req.addressID)
	bsonFilters := bson.D{}
	bsonFilters = append(bsonFilters, bson.E{"_id", bson.D{bson.E{"$eq", addressId}}})
	//bsonFilters = append(bsonFilters, bson.E{"disponible", bson.D{bson.E{"$eq", true}}})
	var address *Address
	err := repo.addressesCollection.FindOne(ctx, bsonFilters, nil).Decode(&address)
	return address, err
}

func (repo *MongoRepository) GetCountries(ctx context.Context, req *Request) ([]*Country, error){
	bsonFilters := bson.M{}
	cur, err := repo.countriesCollection.Find(ctx, bsonFilters, nil)
	var countries []*Country
	for cur.Next(ctx) {
		var country *Country
		if err := cur.Decode(&country); err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}
	return countries, err
}

func (repo *MongoRepository) GetCities(ctx context.Context, req *GetRequest) ([]*City, error){
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
	cur, err := repo.citiesCollection.Find(ctx, bsonFilters, opts)
	var cities []*City
	for cur.Next(ctx) {
		var city *City
		if err := cur.Decode(&city); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, err
}