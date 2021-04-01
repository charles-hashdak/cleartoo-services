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
			fmt.Println(err2)
			return err2
		}
		conversation.Avatar.Url = userRes.User.AvatarUrl
		conversation.Username = userRes.User.Username
		productRes, err3 := s.catalogClient.GetProduct(ctx, &catalogPb.GetRequest{
			UserId: userId,
			ProductId: userId,
		})
		if err3 != nil {
			fmt.Println(err3)
			return err3
		}
		conversation.Product.InCart = productRes.InCart
	}
	res.Conversations = UnmarshalConversations(conversations)
	return nil
}