package realm_services

import (
	"auth-service/internals/domain/models"
	"auth-service/internals/utils/exceptions"
	"fmt"
)

func (u *RealmServices) GetById(id int64) (*models.Realm, *exceptions.AuthError) {

	user, err := u.realmRepo.GetRealmByIdRepo(id)

	if err != nil {
		fmt.Println("cannot get realm by id: ", err)
		return nil, exceptions.NewAuthError(404, "realm not found")
	}

	return user, nil

}
