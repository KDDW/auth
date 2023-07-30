package adapters

import (
	"auth-service/internals/domain/models"
	"context"

	"github.com/uptrace/bun"
)

type BunRealmRepository struct {
	db *bun.DB
}

func NewBunRealmRepository(db *bun.DB) *BunRealmRepository {
	return &BunRealmRepository{
		db: db,
	}
}

func (s *BunRealmRepository) CreateRealmRepo(realm *models.Realm) error {

	ctx := context.Background()
	_, err := s.db.NewInsert().Model(realm).Exec(ctx)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *BunRealmRepository) GetRealmByCodeRepo(code string) (*models.Realm, error) {

	ctx := context.Background()
	realm := new(models.Realm)
	err := s.db.NewSelect().Model(realm).Where("code = ?", code).Scan(ctx)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return realm, nil
}
