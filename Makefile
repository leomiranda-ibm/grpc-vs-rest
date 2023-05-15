rest-server:
	@go run cmd/rest/server/*.go
rest-client:
	@go run cmd/rest/client/*.go --requests=$(REQUESTS)
grpc-server:
	@go run cmd/grpc/server/*.go
grpc-client:
	@go run cmd/grpc/client/*.go --requests=$(REQUESTS)
generate-grpc:
	@protoc --go_out=. --go-grpc_out=. internal/proto/*.proto