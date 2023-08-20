package services

import (
	"auth-service/internals/adapters/repositories"
	auth_services "auth-service/internals/domain/services/auth-services"
	realm_services "auth-service/internals/domain/services/realm-services"
	user_services "auth-service/internals/domain/services/user-services"
)

type Services struct {
	UserServices  *user_services.UserServices
	RealmServices *realm_services.RealmServices
	AuthServices  *auth_services.AuthServices
}

func GetServices(r *repositories.Repositories) *Services {
	userServices := user_services.NewUserServices(r.UserRepo, r.RealmRepo)
	realmServices := realm_services.NewRealmServices(r.RealmRepo)
	authServices := auth_services.NewAuthService(userServices, realmServices)

	return &Services{
		UserServices:  userServices,
		RealmServices: realmServices,
		AuthServices:  authServices,
	}
}
