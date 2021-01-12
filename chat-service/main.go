// chat-service/main.go

package main

import(
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/charles-hashdak/cleartoo-services/chat-service/proto/chat"
	userPb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	"github.com/micro/go-micro/v2"
)

func main(){

	service := micro.NewService(
		micro.Name("cleartoo.chat"),
		micro.Version("latest"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	chatCollection := client.Database("cleartoo").Collection("chat")

	repository := &MongoRepository{chatCollection}

	userClient := userPb.NewUserService("cleartoo.user", service.Client())

	h := &handler{repository, userClient}

	if err := pb.RegisterChatServiceHandler(service.Server(), h); err != nil{
		fmt.Println(err)
	}

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}