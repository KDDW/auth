package realm_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *RealmServices) CreateRealm(dto *dtos.CreateRealmDto) *exceptions.AuthError {

	realmFound, err := u.realmRepo.GetRealmByCodeRepo(dto.Code)

	if err == nil && realmFound != nil {
		msg := fmt.Sprintf("realm %s already exists", dto.Code)
		return exceptions.NewAuthError(409, msg)
	}

	newRealm := &models.Realm{
		Code: dto.Code,
	}

	err = u.realmRepo.CreateRealmRepo(newRealm)

	if err != nil {
		return exceptions.NewAuthError(500, "unknow error while creating a new realm: "+err.Error())
	}

	return nil
}
