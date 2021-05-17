// chat-service/main.go

package main

import(
	"context"
	"fmt"

	pb "github.com/charles-hashdak/cleartoo-services/chat-service/proto/chat"
	userPb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	catalogPb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
)

type handler struct{
	repository
	userClient userPb.UserService
	catalogClient catalogPb.CatalogService
}

func (s *handler) Send(ctx context.Context, chat *pb.Chat, res *pb.SendResponse) error {

	err := s.repository.Send(ctx, MarshalChat(chat))

	if err != nil{
		return nil
	}

	senderRes, err2 := s.userClient.Get(ctx, &userPb.User{
		Id: chat.SenderId,
	})
	if err2 != nil {
		return err2
	}

	_, err3 := s.userClient.SendNotification(ctx, &userPb.Notification{
		UserId: chat.ReceiverId,
		Title: "New message from "+senderRes.User.Username+"!",
		Body: chat.Message,
	})
	if err3 != nil {
		return err3
	}

	res.Sent = true

	return nil
}

func (s *handler) GetChat(ctx context.Context, req *pb.GetChatRequest, res *pb.GetChatResponse) error {
	chats, err := s.repository.GetChat(ctx, MarshalGetChatRequest(req))
	if err != nil {
		return err
	}
	res.Chats = UnmarshalChats(chats)
	return nil
}

func (s *handler) GetConversations(ctx context.Context, req *pb.GetConversationsRequest, res *pb.GetConversationsResponse) error {
	conversations, err := s.repository.GetConversations(ctx, MarshalGetConversationsRequest(req))
	if err != nil {
		return err
	}
	for _, conversation := range conversations {
		fmt.Println(conversation.Product.ID.Hex())
		var userId string
		if conversation.SenderID == req.UserId {
			userId = conversation.ReceiverID
		}else{
			userId = conversation.SenderID
		}
		userRes, err2 := s.userClient.Get(ctx, &userPb.User{
			Id: userId,
		})
		if err2 != nil {
			return err2
		}
		conversation.Avatar.Url = userRes.User.AvatarUrl
		conversation.Username = userRes.User.Username

		productsRes, err4 := s.catalogClient.GetProducts(ctx, &catalogPb.GetRequest{
			UserId: userId,
			Filters: []*catalogPb.Filter{{
				Key: "_id",
				Condition: "$eq",
				Value: conversation.Product.ID.Hex(),
				Hex: true,
			}},
		})
		if err4 != nil {
			return err4
		}

		conversation.Product.InCart = productsRes.Products[0].InCart
	}
	res.Conversations = UnmarshalConversations(conversations)
	return nil
}