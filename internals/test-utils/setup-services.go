package test_utils

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/domain/services"
	"auth-service/internals/dtos"
	"auth-service/internals/infra/repositories"
)

func SetupServices() *services.Services {
	repositories := repositories.GetRepositories(nil)
	return services.GetServices(repositories)
}

var DefaultRealm *dtos.CreateRealmDto = &dtos.CreateRealmDto{
	Code: "test",
}

var DefaultUser *dtos.CreateUserDto = &dtos.CreateUserDto{
	Email:    "test@gmail.com",
	Password: "123456",
	Realm:    "test",
	Name:     "test user",
}

func CreateTestUser(services *services.Services) *models.User {

	services.RealmServices.CreateRealm(DefaultRealm)
	err := services.UserServices.CreateUser(DefaultUser)

	if err != nil {
		panic(err)
	}

	user, _ := services.UserServices.GetByEmailAndRealm(DefaultUser.Email, DefaultUser.Realm)

	return user
}
