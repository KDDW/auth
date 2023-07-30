package main

import (
	"auth-service/internals/adapters"
	"auth-service/internals/domain/services"
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

	userServices := services.NewUserServices(repositories.UserRepo, repositories.RealmRepo)

	err := userServices.CreateUser(&dtos.CreateUserDto{
		Realm:    "app1",
		Password: "123456",
		Email:    "teste2@gmail.com",
	})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("user successfully created")
	}
}
