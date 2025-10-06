.PHONY: run mock

run:
	go run . server

mock:
	go test -v ./internal/mock/...

gen:
	mockery --output=gen/repo --dir=internal/repo --name=IRepository
