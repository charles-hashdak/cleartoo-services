// chat-service/main.go

package main

import(
	"context"
	_ "log"
	"time"
	_ "fmt"

	pb "github.com/charles-hashdak/cleartoo-services/chat-service/proto/chat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

type Chat struct{
	ID 				primitive.ObjectID  `bson:"_id,omitempty"`
	SenderID 		string 				`json:"sender_id"`
	ReceiverID 		string 				`json:"receiver_id"`
	Message 		string 				`json:"message"`
	SendAt 			string 				`json:"send_at"`
	Product 		Product 			`json:"product"`
	Order 			Order 				`json:"order"`
}

type Chats []*Chat

type Conversation struct{
	SenderID 		string
	ReceiverID 		string
	Username 		string
	Avatar 			Photo
	LastChat		string
	SendAt			string
	Product 		Product
	Order 			Order
}

type Conversations []*Conversation

type Product struct{
	ID 				primitive.ObjectID
	Available 		bool
	Title 			string
	Price 			int32
	Photo 			Photo
	InCart 			bool
}

type Products []*Product

type Order struct{
	ID 				primitive.ObjectID
	UserID 			string 
	Products 		Products	
	SubTotal 		float32	
	ShippingFees    float32	
	Taxes 			float32
	Total 			float32
	Status 			string 
	ShippingMethod 	string 
	PaymentMethod 	string
	TrackID 		string
	ShippingStatus	string 
}

type Photo struct{
	ID 				primitive.ObjectID 
	Url 			string
	IsMain 			bool
	Height 			int32
	Width 			int32
}

type SendRequest struct {
	Chat 			Chat
}

type SendResponse struct {
	Sent			bool
}

type GetChatRequest struct {
	SenderID 		string
	ReceiverID 		string
	ProductID 		string
	OrderID 		string
}

type GetChatResponse struct {
	Chats 			Chats
}

type GetConversationsRequest struct {
	UserID 			string
}

type GetConversationsResponse struct {
	Conversations 	Conversations
}

func MarshalSendRequest(req *pb.SendRequest) *SendRequest{
	return &SendRequest{
		Chat: 		*MarshalChat(req.Chat),
	}
}

func UnmarshalSendRequest(req *SendRequest) *pb.SendRequest{
	return &pb.SendRequest{
		Chat: 		UnmarshalChat(&req.Chat),
	}
}

func MarshalSendResponse(req *pb.SendResponse) *SendResponse{
	return &SendResponse{
		Sent: 			req.Sent,
	}
}

func UnmarshalSendResponse(req *SendResponse) *pb.SendResponse{
	return &pb.SendResponse{
		Sent: 			req.Sent,
	}
}

func MarshalGetChatRequest(req *pb.GetChatRequest) *GetChatRequest{
	return &GetChatRequest{
		SenderID: 		req.SenderId,
		ReceiverID: 	req.ReceiverId,
		ProductID: 		req.ProductId,
		OrderID: 		req.OrderId,
	}
}

func UnmarshalGetChatRequest(req *GetChatRequest) *pb.GetChatRequest{
	return &pb.GetChatRequest{
		SenderId: 		req.SenderID,
		ReceiverId: 	req.ReceiverID,
		ProductId: 		req.ProductID,
		OrderId: 		req.OrderID,
	}
}

func MarshalGetChatResponse(req *pb.GetChatResponse) *GetChatResponse{
	return &GetChatResponse{
		Chats: 			MarshalChats(req.Chats),
	}
}

func UnmarshalGetChatResponse(req *GetChatResponse) *pb.GetChatResponse{
	return &pb.GetChatResponse{
		Chats: 			UnmarshalChats(req.Chats),
	}
}

func MarshalGetConversationsRequest(req *pb.GetConversationsRequest) *GetConversationsRequest{
	return &GetConversationsRequest{
		UserID: 		req.UserId,
	}
}

func UnmarshalGetConversationsRequest(req *GetConversationsRequest) *pb.GetConversationsRequest{
	return &pb.GetConversationsRequest{
		UserId: 		req.UserID,
	}
}

func MarshalGetConversationsResponse(req *pb.GetConversationsResponse) *GetConversationsResponse{
	return &GetConversationsResponse{
		Conversations: 			MarshalConversations(req.Conversations),
	}
}

func UnmarshalGetConversationsResponse(req *GetConversationsResponse) *pb.GetConversationsResponse{
	return &pb.GetConversationsResponse{
		Conversations: 			UnmarshalConversations(req.Conversations),
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
	}
}

func UnmarshalProduct(product *Product) *pb.Product{
	return &pb.Product{
		Id:				product.ID.Hex(),
		Available:		product.Available,
		Title:			product.Title,
		Price:			product.Price,
		Photo:			UnmarshalPhoto(&product.Photo),
		InCart:			product.InCart,
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
		SubTotal:		order.SubTotal,
		ShippingFees:	order.ShippingFees,
		Taxes:			order.Taxes,
		Total:			order.Total,
		Status:			order.Status,
		ShippingMethod:	order.ShippingMethod,
		PaymentMethod:	order.PaymentMethod,
		TrackID:		order.TrackId,
		ShippingStatus:	order.ShippingStatus,
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
		TrackId:		order.TrackID,
		ShippingStatus:	order.ShippingStatus,
	}
}

func MarshalChat(chat *pb.Chat) *Chat{
	objId, _ := primitive.ObjectIDFromHex(chat.Id)
	return &Chat{
		ID:				objId,
		SenderID:		chat.SenderId,
		ReceiverID:		chat.ReceiverId,
		Message:		chat.Message,
		SendAt:			chat.SendAt,
		Product:		*MarshalProduct(chat.Product),
		Order:			*MarshalOrder(chat.Order),
	}
}

func UnmarshalChat(chat *Chat) *pb.Chat{
	return &pb.Chat{
		Id:				chat.ID.Hex(),
		SenderId:		chat.SenderID,
		ReceiverId:		chat.ReceiverID,
		Message:		chat.Message,
		SendAt:			chat.SendAt,
		Product:		UnmarshalProduct(&chat.Product),
		Order:			UnmarshalOrder(&chat.Order),
	}
}

func MarshalChats(chats []*pb.Chat) Chats {
	collection := make(Chats, 0)
	for _, chat := range chats {
		collection = append(collection, MarshalChat(chat))
	}
	return collection
}

func UnmarshalChats(chats Chats) []*pb.Chat {
	collection := make([]*pb.Chat, 0)
	for _, chat := range chats {
		collection = append(collection, UnmarshalChat(chat))
	}
	return collection
}

func MarshalConversation(conversation *pb.Conversation) *Conversation{
	return &Conversation{
		SenderID:		conversation.SenderId,
		ReceiverID:		conversation.ReceiverId,
		Username:		conversation.Username,
		Avatar:			*MarshalPhoto(conversation.Avatar),
		LastChat:		conversation.LastChat,
		SendAt:			conversation.SendAt,
		Product:		*MarshalProduct(conversation.Product),
		Order:			*MarshalOrder(conversation.Order),
	}
}

func UnmarshalConversation(conversation *Conversation) *pb.Conversation{
	return &pb.Conversation{
		SenderId:		conversation.SenderID,
		ReceiverId:		conversation.ReceiverID,
		Username:		conversation.Username,
		Avatar:			UnmarshalPhoto(&conversation.Avatar),
		LastChat:		conversation.LastChat,
		SendAt:			conversation.SendAt,
		Product:		UnmarshalProduct(&conversation.Product),
		Order:			UnmarshalOrder(&conversation.Order),
	}
}

func MarshalConversations(conversations []*pb.Conversation) Conversations {
	collection := make(Conversations, 0)
	for _, conversation := range conversations {
		collection = append(collection, MarshalConversation(conversation))
	}
	return collection
}

func UnmarshalConversations(conversations Conversations) []*pb.Conversation {
	collection := make([]*pb.Conversation, 0)
	for _, conversation := range conversations {
		collection = append(collection, UnmarshalConversation(conversation))
	}
	return collection
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
	Send(ctx context.Context, chat *Chat) error
	GetChat(ctx context.Context, req *GetChatRequest) ([]*Chat, error)
	GetConversations(ctx context.Context, req *GetConversationsRequest) ([]*Conversation, error)
}

type MongoRepository struct{
	chatCollection 	*mongo.Collection
}

func (repo *MongoRepository) Send(ctx context.Context, chat *Chat) error{
	chat.SendAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := repo.chatCollection.InsertOne(ctx, chat)
	return err
}

func (repo *MongoRepository) GetChat(ctx context.Context, req *GetChatRequest) ([]*Chat, error){
	var matchStage bson.D
	if req.ProductID != "" {
		productId, _ := primitive.ObjectIDFromHex(req.ProductID)
		matchStage = bson.D{
			{"$match", bson.D{
				{"$or", bson.A{
					bson.D{
						{"senderid", req.SenderID},
						{"receiverid", req.ReceiverID},
					}, 
					bson.D{
						{"receiverid", req.SenderID},
						{"senderid", req.ReceiverID},
					},
				}},
				{"product.id", productId},
			}},
		}
	}
	if req.OrderID != "" {
		orderId, _ := primitive.ObjectIDFromHex(req.OrderID)
		matchStage = bson.D{
			{"$match", bson.D{
				{"$or", bson.A{
					bson.D{
						{"senderid", req.SenderID},
						{"receiverid", req.ReceiverID},
					}, 
					bson.D{
						{"receiverid", req.SenderID},
						{"senderid", req.ReceiverID},
					},
				}},
				{"order.id", orderId},
			}},
		}
	}
	projectStage := bson.D{{"$project", bson.D{{"message", "$message"}, {"receiverid", "$receiverid"}, {"senderid", "$senderid"}, {"sendat", "$sendat"}, {"product", "$product"}, {"order", "$order"}}}}

	cur, err := repo.chatCollection.Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
	var chats []*Chat
	for cur.Next(ctx) {
		var chat *Chat
		if err := cur.Decode(&chat); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, err
}

func (repo *MongoRepository) GetConversations(ctx context.Context, req *GetConversationsRequest) ([]*Conversation, error){
	matchStage := bson.D{{"$match", bson.D{{"$or", bson.A{bson.D{{"senderid", req.UserID}}, bson.D{{"receiverid", req.UserID}}}}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"receiverid", "$receiverid"}, {"senderid", "$senderid"}, {"productid", "$product.id"}, {"orderid", "$order.id"}}}, {"receiverid", bson.D{{"$last", "$receiverid"}}}, {"senderid", bson.D{{"$last", "$senderid"}}}, {"sendat", bson.D{{"$max", "$sendat"}}}, {"lastchat", bson.D{{"$last", "$message"}}}, {"product", bson.D{{"$last", "$product"}}}, {"order", bson.D{{"$last", "$order"}}}}}}

	cur, err := repo.chatCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		return nil, err
	}
	var conversations []*Conversation
	for cur.Next(ctx) {
		var conversation *Conversation
		if err := cur.Decode(&conversation); err != nil {
			return nil, err
		}
		conversations = append(conversations, conversation)
	}
	return conversations, err
}