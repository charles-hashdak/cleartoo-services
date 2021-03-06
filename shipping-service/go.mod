module github.com/charles-hashdak/cleartoo-services/shipping-service

go 1.14

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.6.3 // indirect
	go.mongodb.org/mongo-driver v1.5.0
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
	google.golang.org/genproto v0.0.0-20200914193844-75d14daec038 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)
