package memory_repo

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/domain/ports"
	"auth-service/internals/dtos"
	"errors"
	"strconv"
	"time"
)

type MemoryUserRepository struct {
	db        []models.User
	realmRepo ports.RealmRepository
}

var userId int64 = 1

func NewMemoryUserRepository(realmRepo ports.RealmRepository) *MemoryUserRepository {
	return &MemoryUserRepository{
		db:        make([]models.User, 0),
		realmRepo: realmRepo,
	}
}

func (r *MemoryUserRepository) CreateUserRepo(user *models.User) error {

	realm, err := r.realmRepo.GetRealmByIdRepo(user.RealmID)

	if err != nil {
		return err
	}

	user.ID = userId
	userId++
	user.Realm = realm
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	// hash password is a heavy operation (~ 700ms)
	user.HashPassword()

	r.db = append(r.db, *user)

	return nil
}

func (r *MemoryUserRepository) GetByEmailAndRealmUserRepo(email string, realmID int64) (*models.User, error) {

	for _, user := range r.db {
		if email != "" && realmID != 0 {
			if user.Email == email && user.RealmID == realmID {
				return &user, nil
			}
		}
	}
	return nil, errors.New("user not found")
}

func (r *MemoryUserRepository) GetByIdUserRepo(id int64) (*models.User, error) {
	for _, user := range r.db {
		if id != 0 && user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MemoryUserRepository) DeleteUserRepo(id int64) (int64, error) {

	filtered := []models.User{}
	var affectedRows int64 = 0

	for _, user := range r.db {
		if user.ID != id {
			filtered = append(filtered, user)
		} else {
			affectedRows++
		}
	}

	r.db = filtered

	if affectedRows > 0 {
		return affectedRows, nil
	}

	return 0, errors.New("user not found")
}

func (r *MemoryUserRepository) UpdateUserRepo(id int64, dto *dtos.UpdateUserDto) (int64, error) {

	userFound, err := r.GetByIdUserRepo(id)

	if err != nil {
		return 0, err
	}

	if dto.Email != "" {
		userFound.Email = dto.Email
	}

	if dto.Password != "" {
		userFound.Password = dto.Password
		userFound.HashPassword()
	}

	return 1, nil
}

func (r *MemoryUserRepository) ListUsersRepo(dto *dtos.ListUsersDto) ([]models.User, error) {

	out := []models.User{}

	var recordCount int = 0
	var pageCount int = 1

	recordsPerPage, err := strconv.Atoi(dto.RecordsPerPage)

	if err != nil {
		recordsPerPage = 10
	}

	if recordsPerPage <= 0 {
		recordsPerPage = 10
	}

	page, err := strconv.Atoi(dto.Page)

	if err != nil {
		page = 1
	}

	if page <= 0 {
		page = 1
	}

	for _, user := range r.db {
		if dto.Email == "" && dto.RealmCode == "" {
			out = append(out, user)
			recordCount++
		}

		if dto.Email != "" && dto.RealmCode == "" && user.Email == dto.Email {
			out = append(out, user)
			recordCount++
		}

		if dto.Email == "" && dto.RealmCode != "" && user.Realm.Code == dto.RealmCode {
			out = append(out, user)
			recordCount++
		}
		if recordsPerPage >= recordCount && page == pageCount {
			break
		}
		if recordsPerPage >= recordCount && page < pageCount {
			recordsPerPage = 0
			pageCount++
		}
	}

	return out, nil
}
