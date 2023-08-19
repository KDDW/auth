package controllers

import (
	"auth-service/internals/domain/services"
)

type Controllers struct {
	UserController  *UserController
	RealmController *RealmController
}

func GetControllers(s *services.Services) *Controllers {

	return &Controllers{
		UserController:  NewUserController(s.UserServices),
		RealmController: NewRealmController(s.RealmServices),
	}
}
