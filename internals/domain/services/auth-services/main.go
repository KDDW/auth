package auth_services

import (
	realm_services "auth-service/internals/domain/services/realm-services"
	user_services "auth-service/internals/domain/services/user-services"
)

type AuthServices struct {
	RealmServices *realm_services.RealmServices
	UserServices  *user_services.UserServices
}

func NewAuthService(userServices *user_services.UserServices, RealmServices *realm_services.RealmServices) *AuthServices {
	return &AuthServices{UserServices: userServices, RealmServices: RealmServices}
}
