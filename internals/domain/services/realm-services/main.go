package realm_services

import "auth-service/internals/domain/ports"

type RealmServices struct {
	realmRepo ports.RealmRepository
}

func NewRealmServices(realmRepo ports.RealmRepository) *RealmServices {
	return &RealmServices{
		realmRepo: realmRepo,
	}
}
