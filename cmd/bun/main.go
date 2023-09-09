package main

import (
	"auth-service/cmd/bun/migrations"
	"auth-service/internals/infra/config"
	"auth-service/internals/infra/db"
)

func main() {
	config.LoadEnv()
	db := db.NewDB()
	migrations.Migrate(db)
}
