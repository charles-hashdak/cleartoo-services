build:
	protoc --proto_path=. --go_out=. --micro_out=. proto/shipping/shipping.proto
	sudo docker build -t shipping-service .

run:
	sudo docker run --net="host" \
		-p 50051 \
	    -e DB_HOST=localhost \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		shipping-service