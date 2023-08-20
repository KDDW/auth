package user_services

import (
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *UserServices) DeleteUser(id int64) *exceptions.AuthError {

	userFound, err := u.userRepo.GetByIdUserRepo(id)

	if err != nil {
		msg := fmt.Sprintf("cannot found user with id %d:  %s", id, err.Error())
		fmt.Println(msg)
		return exceptions.NewAuthError(404, "user not found")
	}

	_, err = u.userRepo.DeleteUserRepo(userFound.ID)

	if err != nil {
		fmt.Println("cannot delete user: ", err)
		return exceptions.NewAuthError(500, "cannot delete user")
	}

	return nil

}
