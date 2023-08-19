package realm_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *RealmServices) GetRealmByCode(code string) (*models.Realm, *exceptions.AuthError) {

	realmFound, err := u.realmRepo.GetRealmByCodeRepo(code)

	if err != nil {
		fmt.Println("err != nil on GetRealmByCode service: ", err.Error())
		return nil, exceptions.NewAuthError(404, "realm not found")
	}

	return realmFound, nil
}
