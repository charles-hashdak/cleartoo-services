// chat-service/main.go

package main

import(
	"context"
	"fmt"
	"strconv"
	"time"

	pb "github.com/charles-hashdak/cleartoo-services/chat-service/proto/chat"
	userPb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	catalogPb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	orderPb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
)

type handler struct{
	repository
	userClient userPb.UserService
	catalogClient catalogPb.CatalogService
	orderClient orderPb.OrderService
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
	var convs []*Conversation
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

		if conversation.Product.ID.Hex() != "000000000000000000000000" {
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
			if len(productsRes.Products) > 0 {
				conversation.Product.InCart = productsRes.Products[0].InCart
				if len(productsRes.Products[0].Offers) > 0 {
					layout := "2006-01-02 15:04:05"
					lastOfferUpdatedAt, _ := time.Parse(layout, productsRes.Products[0].Offers[len(productsRes.Products[0].Offers) - 1].UpdatedAt)
					lastChatSendAt, _ := time.Parse(layout, conversation.SendAt)
					if lastOfferUpdatedAt.After(lastChatSendAt) {
						conversation.SendAt = productsRes.Products[0].Offers[len(productsRes.Products[0].Offers) - 1].UpdatedAt
						conversation.LastChat = "฿"+strconv.Itoa(int(productsRes.Products[0].Offers[len(productsRes.Products[0].Offers) - 1].Amount))
					}
				}
				convs = append(convs, conversation)
			}
		}

		if conversation.Order.ID.Hex() != "000000000000000000000000" {
			orderRes, err4 := s.orderClient.GetSingleOrder(ctx, &orderPb.GetSingleRequest{
				OrderId: conversation.Order.ID.Hex(),
			})
			if err4 != nil {
				return err4
			}
			if orderRes.Order != nil {
				if len(orderRes.Order.Offers) > 0 {
					layout := "2006-01-02 15:04:05"
					lastOfferUpdatedAt, _ := time.Parse(layout, orderRes.Order.Offers[len(orderRes.Order.Offers) - 1].UpdatedAt)
					lastChatSendAt, _ := time.Parse(layout, conversation.SendAt)
					if lastOfferUpdatedAt.After(lastChatSendAt) {
						conversation.SendAt = orderRes.Order.Offers[len(orderRes.Order.Offers) - 1].UpdatedAt
						conversation.LastChat = "฿"+strconv.Itoa(int(orderRes.Order.Offers[len(orderRes.Order.Offers) - 1].Amount))
					}
				}
				convs = append(convs, conversation)
			}
		}
	}
	ownProductsRes, err := s.catalogClient.GetProducts(ctx, &catalogPb.GetRequest{
		UserId: userId,
		Filters: []*catalogPb.Filter{{
			Key: "owner.ownerid",
			Condition: "$eq",
			Value: userId,
		}},
	})
	if err != nil {
		return err
	}
	for _, product := ownProductsRes.Products {
		if len(product.Offers) > 0 {
			for _, offer := range product.Offers {
				if len(convs.filter(function(conv){ return conv.Product.ID.Hex() == product.Id && (conv.SenderID == offer.UserID || conv.ReceiverID == offer.UserID) })) == 0 {
					conversation := &Conversation{
						SenderID: offer.UserID,
						ReceiverID: userId,
						LastChat: "฿"+strconv.Itoa(int(offer.Amount)),
						SendAt:	offer.UpdatedAt,
						Product: product,
					}
					convs = append(convs, conversation)
				}
			}
		}
	}
	offersProductsRes, err := s.catalogClient.GetProducts(ctx, &catalogPb.GetRequest{
		UserId: userId,
		Filters: []*catalogPb.Filter{{
			Key: "offers.userid",
			Condition: "$elemMatch",
			Value: userId,
		}},
	})
	if err != nil {
		return err
	}
	for _, product := offersProductsRes.Products {
		if len(convs.filter(function(conv){ return conv.Product.Id.Hex() == product.Id })) == 0 && len(product.Offers) > 0 {
			offer := product.Offers[len(product.Offers) - 1]
			conversation := &Conversation{
				SenderID: userId,
				ReceiverID: product.Owner.OwnerId,
				LastChat: "฿"+strconv.Itoa(int(offer.Amount)),
				SendAt:	offer.UpdatedAt,
				Product: product,
			}
			convs = append(convs, conversation)
		}
	}
	layout := "2006-01-02 15:04:05"
	sort.SliceStable(convs, func(i, j int) bool {
		prevSendAt, _ := time.Parse(layout, convs[i].SendAt)
		nextSendAt, _ := time.Parse(layout, convs[j].SendAt)
	    return prevSendAt.Before(nextSendAt)
	})
	res.Conversations = UnmarshalConversations(convs)
	return nil
}