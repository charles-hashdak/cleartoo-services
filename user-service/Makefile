build:
	protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto
	sudo docker build -t user-service .

run:
	sudo docker run --net="host" \
		-p 50051 \
	    -e DB_NAME=cleartoo \
	    -e DB_HOST=localhost \
	    -e DB_PASSWORD=Postgres784512! \
	    -e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-service