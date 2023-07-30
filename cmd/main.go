package main

import (
	"auth-service/internals/adapters"
	"auth-service/internals/domain/services"
	"auth-service/internals/infra/config"
	"auth-service/internals/infra/db"
	"auth-service/internals/infra/db/migrations"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	config.LoadEnv()
	db := db.NewDB()

	migrations.Migrate(db)

	repositories := adapters.GetRepositories(db)

	userServices := services.NewUserServices(repositories.UserRepo, repositories.RealmRepo)

	user, err := userServices.GetByEmailAndRealm("teste2@gmail.com", "app1")

	if err != nil {
		log.Fatal(err)
	}

	jsonUser, err := json.Marshal(user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", string(jsonUser))
}
