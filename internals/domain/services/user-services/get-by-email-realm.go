package user_services

import (
	"auth-service/internals/domain/models"
	"errors"
	"fmt"
)

func (u *UserServices) GetByEmailAndRealm(email string, realm string) (*models.User, error) {

	realmFound, err := u.realmRepo.GetRealmByCodeRepo(realm)

	if err != nil {
		fmt.Println("cannot get realm: ", err)
		return nil, errors.New("realm not found")
	}

	user, err := u.userRepo.GetByEmailAndRealmUserRepo(email, realmFound.ID)

	if err != nil {
		fmt.Println("cannot get user by email and realm: ", err)
		return nil, errors.New("user not found")
	}

	return user, nil

}
