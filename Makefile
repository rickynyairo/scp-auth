# Makefile for the user microservice

build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/auth/user.proto
	@echo 'Build Docker image'
	docker build -t shippy-auth .

run:
	docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=localhost \
		-e DB_NAME=postgres \
		-e DB_USER=postgres \
		-e DB_PASSWORD="" \
		-e DB_PORT=5432 \
		shippy-auth:latest

run-go:
	go run main.go repository.go token_service.go handler.go database.go