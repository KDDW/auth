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

func (u *UserServices) DeleteUser(id int64) error {

	_, err := u.userRepo.DeleteUserRepo(id)

	if err != nil {
		fmt.Println("cannot delete user: ", err)
		return errors.New("cannot delete user")
	}

	return nil

}

func (u *UserServices) GetById(id int64) (*models.User, error) {

	user, err := u.userRepo.GetByIdUserRepo(id)

	if err != nil {
		fmt.Println("cannot get user by id: ", err)
		return nil, errors.New("cannot get user")
	}

	return user, nil

}

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
