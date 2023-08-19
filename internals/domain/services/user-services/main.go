package user_services

import "auth-service/internals/domain/ports"

type UserServices struct {
	userRepo  ports.UserRepository
	realmRepo ports.RealmRepository
}

func NewUserServices(userRepo ports.UserRepository, realmRepo ports.RealmRepository) *UserServices {
	return &UserServices{
		userRepo:  userRepo,
		realmRepo: realmRepo,
	}

}
