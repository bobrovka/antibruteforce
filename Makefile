gen.proto:
	protoc --go_out=plugins=grpc:pkg/antibruteforce api/api.proto

build:
	go build -o bin ./cmd/main

redis.run:
	docker run --name redis-antibruteforce -p 6379:6379 -d redis
