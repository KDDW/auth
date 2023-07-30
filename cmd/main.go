package main
import (
	"auth-service/internals/adapters"
	"auth-service/internals/dtos"
	"auth-service/internals/infra/config"
	"auth-service/internals/infra/db"
	"auth-service/internals/infra/db/migrations"
	"fmt"
	"log"
)

func main() {
	config.LoadEnv()
	db := db.NewDB()

	migrations.Migrate(db)
	repositories := adapters.GetRepositories(db)

