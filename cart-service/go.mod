module github.com/charles-hashdak/cleartoo-services/cart-service

go 1.14

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/miekg/dns v1.1.31 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	go.mongodb.org/mongo-driver v1.5.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200915084602-288bc346aa39 // indirect
	golang.org/x/tools v0.0.0-20200915173823-2db8f0ff891c // indirect
	google.golang.org/genproto v0.0.0-20200914193844-75d14daec038 // indirect
	google.golang.org/grpc v1.30.0 // indirect
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
)
