package adapters

import (
	"auth-service/internals/domain/ports"

	"github.com/uptrace/bun"
)

type Repositories struct {
	UserRepo  ports.UserRepository
	RealmRepo ports.RealmRepository
}

func GetRepositories(db *bun.DB) *Repositories {

	return &Repositories{
		UserRepo:  NewBunUserRepository(db),
		RealmRepo: NewBunRealmRepository(db),
	}
}
