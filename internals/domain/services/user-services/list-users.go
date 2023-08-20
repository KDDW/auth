package user_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *UserServices) ListUsers(dto *dtos.ListUsersDto) ([]models.User, *exceptions.AuthError) {

	users, err := u.userRepo.ListUsersRepo(dto)

	if err != nil {
		fmt.Println("cannot list users: " + err.Error())
		return nil, exceptions.NewAuthError(500, "unknow error while listing users")
	}

	return users, nil
}
