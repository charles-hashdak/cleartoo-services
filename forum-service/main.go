// forum-service/main.go

package main

import(
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/charles-hashdak/cleartoo-services/forum-service/proto/forum"
	userPb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	"github.com/micro/go-micro/v2"
)

func main(){

	service := micro.NewService(
		micro.Name("cleartoo.forum"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	subjectCollection := client.Database("cleartoo").Collection("subject")
	commentCollection := client.Database("cleartoo").Collection("comment")

	repository := &MongoRepository{subjectCollection, commentCollection}

	userClient := userPb.NewUserService("cleartoo.user", service.Client())

	h := &handler{repository, userClient}

	if err := pb.RegisterForumServiceHandler(service.Server(), h); err != nil{
		fmt.Println(err)
	}

	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}