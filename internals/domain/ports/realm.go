package ports

import (
	"auth-service/internals/domain/models"
)

type RealmRepository interface {
	CreateRealmRepo(realm *models.Realm) error
	GetRealmByCodeRepo(code string) (*models.Realm, error)
}
