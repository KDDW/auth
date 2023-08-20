package ports

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
)

type UserRepository interface {
	CreateUserRepo(user *models.User) error
	UpdateUserRepo(id int64, dto *dtos.UpdateUserDto) (int64, error)
	DeleteUserRepo(id int64) (int64, error)
	GetByIdUserRepo(id int64) (*models.User, error)
	GetByEmailAndRealmUserRepo(email string, realmID int64) (*models.User, error)
	ListUsersRepo(dto *dtos.ListUsersDto) ([]models.User, error)
}
