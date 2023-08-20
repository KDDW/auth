package repositories

import (
	"auth-service/internals/domain/ports"
	bun_repo "auth-service/internals/infra/repositories/bun-repo"
	memory_repo "auth-service/internals/infra/repositories/memory-repo"
	"os"

	"github.com/uptrace/bun"
)

type Repositories struct {
	UserRepo  ports.UserRepository
	RealmRepo ports.RealmRepository
}

func GetRepositories(db *bun.DB) *Repositories {

	var realmRepo ports.RealmRepository
	var userRepo ports.UserRepository

	dbType := os.Getenv("DB_TYPE")

	if dbType == "" && db != nil {
		dbType = "bun"
	}

	if db == nil {
		dbType = "memory"
	}

	if dbType == "bun" {
		realmRepo = bun_repo.NewBunRealmRepository(db)
		userRepo = bun_repo.NewBunUserRepository(db)
	} else {
		realmRepo = memory_repo.NewMemoryRealmRepository()
		userRepo = memory_repo.NewMemoryUserRepository(realmRepo)
	}

	return &Repositories{
		UserRepo:  userRepo,
		RealmRepo: realmRepo,
	}
}
