package user_services

import (
	"auth-service/internals/dtos"
	"errors"
	"fmt"
)

func (u *UserServices) UpdateUser(id int64, dto *dtos.UpdateUserDto) error {

	userFound, err := u.userRepo.GetByIdUserRepo(id)

	if err != nil {
		return errors.New("User not found")
	}

	affectedRows, err := u.userRepo.UpdateUserRepo(userFound.ID, dto)

	if err != nil {
		fmt.Println("cannot updater user: ", err)
		return errors.New("cannot update user")
	}

	if affectedRows == 1 {
		fmt.Println("user successfully updated")
	}

	return nil
}
