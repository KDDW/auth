package repositories

import (
	bun_repo "auth-service/internals/adapters/repositories/bun-repo"
	"auth-service/internals/domain/ports"

	"github.com/uptrace/bun"
)

type Repositories struct {
	UserRepo  ports.UserRepository
	RealmRepo ports.RealmRepository
}

func GetRepositories(db *bun.DB) *Repositories {

	return &Repositories{
		UserRepo:  bun_repo.NewBunUserRepository(db),
		RealmRepo: bun_repo.NewBunRealmRepository(db),
	}
}
