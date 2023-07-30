package services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/domain/ports"
	"auth-service/internals/dtos"
	"errors"
	"fmt"
)

type UserServices struct {
	userRepo  ports.UserRepository
	realmRepo ports.RealmRepository
}

func NewUserServices(userRepo ports.UserRepository, realmRepo ports.RealmRepository) *UserServices {
	return &UserServices{
		userRepo:  userRepo,
		realmRepo: realmRepo,
	}

}

func (u *UserServices) CreateUser(dto *dtos.CreateUserDto) error {

	realmFound, err := u.realmRepo.GetRealmByCodeRepo(dto.Realm)

	if err != nil {
		return errors.New("realm not found")
	}

	userFound, err := u.userRepo.GetByEmailAndRealmUserRepo(dto.Email, realmFound.ID)

	if userFound != nil && err == nil {
		return errors.New("user already registered")
	}

	newUser := models.User{
		Email:    dto.Email,
		RealmID:  realmFound.ID,
		Password: dto.Password,
	}

	err = u.userRepo.CreateUserRepo(&newUser)

	if err != nil {
		fmt.Println("cannot create user: ", err)
		return errors.New("cannot create user")
	}

	return nil
}

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
