module github.com/charles-hashdak/cleartoo-services/order-service

go 1.14

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/charles-hashdak/cleartoo-services/cart-service v0.0.0-20210310052804-c83c37d687af
	github.com/charles-hashdak/cleartoo-services/user-service v0.0.0-20210419163739-5fdca6b28463 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
	github.com/spf13/viper v1.6.3 // indirect
	go.mongodb.org/mongo-driver v1.5.0
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)
