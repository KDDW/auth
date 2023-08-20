package memory_repo

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
	"errors"
	"time"
)

type MemoryRealmRepository struct {
	db []models.Realm
}

var realmId int64 = 1

func NewMemoryRealmRepository() *MemoryRealmRepository {
	return &MemoryRealmRepository{
		db: make([]models.Realm, 0),
	}
}

func (m *MemoryRealmRepository) CreateRealmRepo(realm *models.Realm) error {
	realm.ID = realmId
	realmId++
	now := time.Now()
	realm.CreatedAt = now
	realm.UpdatedAt = now
	m.db = append(m.db, *realm)
	return nil
}

func (m *MemoryRealmRepository) GetRealmByCodeRepo(code string) (*models.Realm, error) {

	for _, realm := range m.db {
		if realm.Code == code {
			return &realm, nil
		}
	}

	return nil, errors.New("realm not found")
}

func (m *MemoryRealmRepository) ListRealmsRepo(dto *dtos.ListRealmsDto) ([]models.Realm, error) {

	out := []models.Realm{}

	if dto.Code == "" {
		out = append(out, m.db...)
		return out, nil
	}

	for _, realm := range m.db {
		if realm.Code == dto.Code {
			out = append(out, realm)
		}
	}

	return out, nil

}

func (m *MemoryRealmRepository) GetRealmByIdRepo(id int64) (*models.Realm, error) {

	for _, realm := range m.db {
		if realm.ID == id {
			return &realm, nil
		}
	}

	return nil, errors.New("realm not found")
}
