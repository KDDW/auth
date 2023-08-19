package user_services

import (
	"errors"
	"fmt"
)

func (u *UserServices) DeleteUser(id int64) error {

	_, err := u.userRepo.DeleteUserRepo(id)

	if err != nil {
		fmt.Println("cannot delete user: ", err)
		return errors.New("cannot delete user")
	}

	return nil

}
