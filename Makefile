proto:
	protoc internal/**/pb/*.proto --go_out=.
	protoc internal/**/pb/*.proto --go-grpc_out=.
	go mod tidy

swagger:
	swag init -g cmd/main.go

server:
	go run cmd/main.go