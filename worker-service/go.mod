module github.com/charles-hashdak/cleartoo-services/worker-service

go 1.14

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)
