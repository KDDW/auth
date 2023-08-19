package realm_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
)

func (u *RealmServices) ListRealms(dto *dtos.ListRealmsDto) ([]models.Realm, *exceptions.AuthError) {

	realms, err := u.realmRepo.ListRealmsRepo(dto)

	if err != nil {
		msg := "can't list realms: " + err.Error()
		return realms, exceptions.NewAuthError(500, msg)
	}

	return realms, nil
}
