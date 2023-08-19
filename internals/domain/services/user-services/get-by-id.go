package user_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *UserServices) GetById(id int64) (*models.User, *exceptions.AuthError) {

	user, err := u.userRepo.GetByIdUserRepo(id)

	if err != nil {
		fmt.Println("cannot get user by id: ", err)
		return nil, exceptions.NewAuthError(404, "user not found")
	}

	return user, nil

}
