package user_services

import (
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *UserServices) UpdateUser(id int64, dto *dtos.UpdateUserDto) *exceptions.AuthError {

	userFound, err := u.userRepo.GetByIdUserRepo(id)

	if err != nil {
		fmt.Println("cannot get user: ", err.Error())
		return exceptions.NewAuthError(404, "user not found")
	}

	affectedRows, err := u.userRepo.UpdateUserRepo(userFound.ID, dto)

	if err != nil {
		fmt.Println("cannot update user: ", err.Error())
		return exceptions.NewAuthError(500, "cannot update user")
	}

	if affectedRows == 1 {
		fmt.Println("user successfully updated")
	}

	return nil
}
