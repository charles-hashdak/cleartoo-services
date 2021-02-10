module github.com/charles-hashdak/cleartoo-services/user-service

go 1.14

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/miekg/dns v1.1.31 // indirect
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	google.golang.org/grpc v1.30.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	google.golang.org/protobuf v1.25.0
)
