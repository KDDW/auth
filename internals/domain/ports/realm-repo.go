package ports

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
)

type RealmRepository interface {
	CreateRealmRepo(realm *models.Realm) error
	GetRealmByCodeRepo(code string) (*models.Realm, error)
	ListRealmsRepo(dto *dtos.ListRealmsDto) ([]models.Realm, error)
}
