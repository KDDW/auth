package main

import (
	"auth-service/cmd/bun/migrations"
	"auth-service/internals/adapters/controllers"
	"auth-service/internals/domain/services"
	"auth-service/internals/infra/config"
	"auth-service/internals/infra/db"
	httpserver "auth-service/internals/infra/http-server"
	"auth-service/internals/infra/repositories"
)

func main() {
	config.LoadEnv()

	db := db.NewDB()
	migrations.Migrate(db)

	repositories := repositories.GetRepositories(db)
	services := services.GetServices(repositories)
	controllers := controllers.GetControllers(services)

	server := httpserver.CreateServer()

	httpserver.RegisterControllers(server, controllers)
	httpserver.Listen(server)
}
