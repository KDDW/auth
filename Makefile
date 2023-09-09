
FIBER_MAIN=cmd/fiber/main.go
MIGRATION_MAIN=cmd/bun/main.go

all: run

run:
	@go run $(FIBER_MAIN)

migrate:
	@go run $(MIGRATION_MAIN) migrate

create-migration:
	@go run scripts/create-migration-file.go

build:
	@go build -o bin/auth $(FIBER_MAIN)

rdb:
	@docker compose rm -f
	@docker rm -f auth-db-1
	@docker system prune -f
	@docker compose up -d

docker-rm-api:
	@docker rm -f auth-api-1 && docker rmi -f auth-api


test:
	@go test ./... | grep -v "no test files"
