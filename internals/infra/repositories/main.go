package repositories

import (
	"auth-service/internals/domain/ports"
	bun_repo "auth-service/internals/infra/repositories/bun-repo"
	memory_repo "auth-service/internals/infra/repositories/memory-repo"
	"os"
	"strings"

	"github.com/uptrace/bun"
)

type Repositories struct {
	UserRepo  ports.UserRepository
	RealmRepo ports.RealmRepository
}

func GetRepositories(db *bun.DB) *Repositories {

	var realmRepo ports.RealmRepository
	var userRepo ports.UserRepository

	if strings.ToLower(os.Getenv("ENV")) == "test" {
		realmRepo = memory_repo.NewMemoryRealmRepository()
		userRepo = memory_repo.NewMemoryUserRepository(realmRepo)
		return &Repositories{
			UserRepo:  userRepo,
			RealmRepo: realmRepo,
		}
	}

	realmRepo = bun_repo.NewBunRealmRepository(db)
	userRepo = bun_repo.NewBunUserRepository(db)

	return &Repositories{
		UserRepo:  userRepo,
		RealmRepo: realmRepo,
	}
}
