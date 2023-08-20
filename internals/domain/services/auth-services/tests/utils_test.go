package auth_services_test

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/domain/services"
	"auth-service/internals/dtos"
)

var defaultRealm *dtos.CreateRealmDto = &dtos.CreateRealmDto{
	Code: "test",
}

var defaultUser *dtos.CreateUserDto = &dtos.CreateUserDto{
	Email:    "test@gmail.com",
	Password: "123456",
	Realm:    "test",
	Name:     "test user",
}

func createTestUser(services *services.Services) *models.User {

	services.RealmServices.CreateRealm(defaultRealm)
	err := services.UserServices.CreateUser(defaultUser)

	if err != nil {
		panic(err)
	}

	user, _ := services.UserServices.GetByEmailAndRealm(defaultUser.Email, defaultUser.Realm)

	return user
}
