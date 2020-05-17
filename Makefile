.PHONY: gen.proto
gen.proto:
	protoc --go_out=plugins=grpc:pkg/antibruteforce api/api.proto

.PHONY: build
build:
	go build -o bin/abf ./cmd/main

.PHONY: run
run:
	docker-compose -f deployment/docker-compose.yaml up --build -d

.PHONY: stop
stop:
	docker-compose -f deployment/docker-compose.yaml down

.PHONY: test
test:
	go test -race -v -count=1 ./...
