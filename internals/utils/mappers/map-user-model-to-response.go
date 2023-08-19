package mappers

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
)

func MapUserModelToResponse(user *models.User) *dtos.GetUserReponse {

	return &dtos.GetUserReponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Realm: &dtos.GetRealmResponse{
			ID:   user.RealmID,
			Code: user.Realm.Code,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
