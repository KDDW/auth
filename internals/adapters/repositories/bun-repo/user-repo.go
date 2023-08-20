package bun_repo

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
	"context"
	"strconv"

	"github.com/uptrace/bun"
)

/*
	implements internals/domain/ports/user-repo
*/

type BunUserRepository struct {
	db *bun.DB
}

func NewBunUserRepository(db *bun.DB) *BunUserRepository {
	return &BunUserRepository{
		db: db,
	}
}

func (s *BunUserRepository) CreateUserRepo(user *models.User) error {

	ctx := context.Background()
	user.HashPassword()
	_, err := s.db.NewInsert().Model(user).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *BunUserRepository) GetByEmailAndRealmUserRepo(email string, realmID int64) (*models.User, error) {

	ctx := context.Background()
	user := new(models.User)

	err := s.db.NewSelect().Model(user).Relation("Realm").Where("realm_id = ?", realmID).Where("email = ?", email).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *BunUserRepository) GetByIdUserRepo(id int64) (*models.User, error) {

	ctx := context.Background()
	user := new(models.User)

	err := s.db.NewSelect().Model(user).Relation("Realm").Where("u.id = ?", id).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *BunUserRepository) DeleteUserRepo(id int64) (int64, error) {

	ctx := context.Background()

	res, err := s.db.NewDelete().Model(&models.User{}).Where("id = ?", id).Exec(ctx)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s *BunUserRepository) UpdateUserRepo(id int64, dto *dtos.UpdateUserDto) (int64, error) {

	ctx := context.Background()

	newUser := models.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	if dto.Email == "" && dto.Password == "" {
		return 0, nil
	}

	query := s.db.NewUpdate().Model(&newUser).Where("id = ?", id)

	if dto.Email != "" {
		query.Set("email = ?", newUser.Email)
	}

	if dto.Password != "" {
		newUser.HashPassword()
		query.Set("password = ?", newUser.Password)
	}

	res, err := query.Exec(ctx)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (u *BunUserRepository) ListUsersRepo(dto *dtos.ListUsersDto) ([]models.User, error) {

	ctx := context.Background()

	users := make([]models.User, 0)

	query := u.db.NewSelect().Model(&users).Relation("Realm")

	if dto.Email != "" {
		query.Where("email = ?", dto.Email)
	}

	if dto.RealmCode != "" {
		query.Where("realm.code = ?", dto.RealmCode)
	}

	var recordsPerPage int
	var page int

	if dto.RecordsPerPage != "" {
		recordsPerPageNbr, err := strconv.Atoi(dto.RecordsPerPage)
		if err != nil {
			return nil, err
		}
		recordsPerPage = recordsPerPageNbr
	}

	if dto.Page != "" {
		pageNbr, err := strconv.Atoi(dto.Page)
		if err != nil {
			return nil, err
		}
		page = pageNbr
	}

	if recordsPerPage <= 0 {
		recordsPerPage = 100
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * recordsPerPage

	query.Limit(recordsPerPage).Offset(offset)

	err := query.Scan(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
