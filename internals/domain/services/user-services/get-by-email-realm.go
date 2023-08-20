package user_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *UserServices) GetByEmailAndRealm(email string, realm string) (*models.User, *exceptions.AuthError) {

	realmFound, err := u.realmRepo.GetRealmByCodeRepo(realm)

	if err != nil {
		fmt.Println("cannot get realm: ", err)
		return nil, exceptions.NewAuthError(404, "realm not found")
	}

	user, err := u.userRepo.GetByEmailAndRealmUserRepo(email, realmFound.ID)

	if err != nil {
		fmt.Println("cannot get user by email and realm: ", err)
		return nil, exceptions.NewAuthError(404, "user not found")
	}

	return user, nil

}
