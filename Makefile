

all: run

run:
	@go run cmd/main.go

migrate:
	@go run cmd/main.go migrate


build:
	@go build -o bin/auth cmd/main.go


rdb:
	@docker compose rm -f
	@docker rm -f auth-db-1
	@docker system prune -f
	@docker compose up -d



test:
	@go test ./... | grep -v "no test files"