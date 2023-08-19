package user_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
)

func (u *UserServices) CreateUser(dto *dtos.CreateUserDto) *exceptions.AuthError {

	realmFound, err := u.realmRepo.GetRealmByCodeRepo(dto.Realm)

	if err != nil {
		return exceptions.NewAuthError(404, "realm not found")
	}

	userFound, err := u.userRepo.GetByEmailAndRealmUserRepo(dto.Email, realmFound.ID)

	if userFound != nil && err == nil {
		return exceptions.NewAuthError(409, "user already registered")
	}

	newUser := models.User{
		Email:    dto.Email,
		RealmID:  realmFound.ID,
		Password: dto.Password,
		Name:     dto.Name,
	}

	err = u.userRepo.CreateUserRepo(&newUser)

	if err != nil {
		return exceptions.NewAuthError(500, "unknow error: "+err.Error())
	}

	return nil
}
