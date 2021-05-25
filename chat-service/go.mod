module github.com/charles-hashdak/cleartoo-services/chat-service

go 1.14

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/charles-hashdak/cleartoo-services/catalog-service v0.0.0-20210428162824-3269f5674c22
	github.com/charles-hashdak/cleartoo-services/user-service v0.0.0-20210428162824-3269f5674c22
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.5.0
	google.golang.org/protobuf v1.25.0
)
