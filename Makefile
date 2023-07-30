

all: run

run:
	@go run cmd/main.go


rdb:
	@docker compose rm -f
	@docker rm -f auth-db-1
	@docker system prune -f
	@docker compose up -d
