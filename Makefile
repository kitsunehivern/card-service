.PHONY: run mock

run:
	go run . server

mock:
	go test -v ./internal/mock/...

mockgen:
	mockery --output=gen/repo --dir=internal/repo --name=IRepository

protogen:
	 protoc --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/card.proto
