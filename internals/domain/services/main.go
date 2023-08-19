package services

import (
	"auth-service/internals/adapters/repositories"
	realm_services "auth-service/internals/domain/services/realm-services"
	user_services "auth-service/internals/domain/services/user-services"
)

type Services struct {
	UserServices  *user_services.UserServices
	RealmServices *realm_services.RealmServices
}

func GetServices(r *repositories.Repositories) *Services {
	return &Services{
		UserServices:  user_services.NewUserServices(r.UserRepo, r.RealmRepo),
		RealmServices: realm_services.NewRealmServices(r.RealmRepo),
	}
}
