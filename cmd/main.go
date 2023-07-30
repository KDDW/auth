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

	realmFound, err := repositories.RealmRepo.GetRealmByCodeRepo("stj")

	if err != nil {
		log.Fatal("cannot get realm", err)
	}

	userFound, err := repositories.UserRepo.GetByEmailAndRealmUserRepo("teste@gmail.com", realmFound.ID)

	if err != nil {
		log.Fatal("user not found", err)
	}

	rowsAffected, err := repositories.UserRepo.UpdateUserRepo(userFound.ID, &dtos.UpdateUserDto{
		Password: "123456",
	})

	if err != nil {
		fmt.Println("cannot update user: ", err)
	}

	if rowsAffected == 1 {
		fmt.Println("User successfully updated")
	}
}
