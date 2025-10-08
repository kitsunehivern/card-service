http-mem:
	go run . server memory

http-psql:
	go run . server postgres

test:
	go clean -testcache && go test -race -v ./internal/test/...

mockgen:
	mockery

protogen:
	 protoc --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/card.proto

docker:
	docker run --name carddb \
    	-e POSTGRES_USER=user \
    	-e POSTGRES_PASSWORD=pass \
    	-e POSTGRES_DB=carddb \
    	-p 5432:5432 \
    	-d postgres:16

db-update:
	atlas migrate diff \
      --dir "file://migration" \
      --to  "file://db" \
      --dev-url "postgres://user:pass@localhost:5432/carddb?sslmode=disable"

db-migrate:
	atlas migrate apply \
      --dir "file://migration" \
      --url "postgres://user:pass@localhost:5432/carddb?sslmode=disable"
