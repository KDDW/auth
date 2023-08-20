package repositories

import (
	"auth-service/internals/domain/ports"
	bun_repo "auth-service/internals/infra/repositories/bun-repo"

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
