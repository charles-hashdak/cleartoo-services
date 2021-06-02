package main

import (
	"log"
	"sync"

	pb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	orderPb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	"github.com/micro/go-micro/v2"
)

const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		username varchar(225) default null not null unique,
		password varchar(225) not null,
		company varchar(125),
		description text,
		rating float(8) not null,
		avatar_url varchar(225),
		cover_url varchar(225),
		followers_count int,
		following_count int,
		age int,
		fb_token varchar(225),
		push_token varchar(225),
		rating_count int,
		apple_user_id varchar(225),
		primary key (id)
	);

	create table if not exists followers (
		id varchar(36) not null,
		follower_id varchar(36) not null,
		user_id varchar(36) not null,
		primary key (id)
	);
	
	create table if not exists ratings (
		id varchar(36) not null,
		rater_id varchar(36) not null,
		user_id varchar(36) not null,
		order_id varchar(36) not null unique,
		rate float(8) not null,
		primary key (id)
	);
`

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Run schema query on start-up, as we're using "create if not exists"
	// this will only be ran once. In order to create updates, you'll need to
	// use a migrations library
	db.AutoMigrate(&pb.User{}, &pb.Follower{}, &pb.Rating{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("cleartoo.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	publisher := micro.NewPublisher("user.created", srv.Client())

	mutex := sync.Mutex{}

	orderClient := orderPb.NewOrderService("cleartoo.order", srv.Client())

	// Register handler
	if err := pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, publisher, mutex, orderClient}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := srv.Run(); err != nil {
		log.Panic(err)
	}
}
