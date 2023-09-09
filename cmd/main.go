package main

import (
	"auth-service/internals/adapters/controllers"
	"auth-service/internals/domain/services"
	"auth-service/internals/infra/config"
	"auth-service/internals/infra/db"
	"auth-service/internals/infra/db/migrations"
	httpserver "auth-service/internals/infra/http-server"
	"auth-service/internals/infra/repositories"
	"os"
)

func main() {
	config.LoadEnv()

	if len(os.Args) == 2 {
		args := os.Args[1:]

		if args[0] == "migrate" {
			db := db.NewDB()
			migrations.Migrate(db)
			return
		}
	}

	db := db.NewDB()
	migrations.Migrate(db)

	repositories := repositories.GetRepositories(db)
	services := services.GetServices(repositories)
	controllers := controllers.GetControllers(services)

	server := httpserver.CreateServer()

	httpserver.RegisterControllers(server, controllers)
	httpserver.Listen(server)
}
