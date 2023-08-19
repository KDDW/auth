package dtos

import (
	"time"
)

type GetUserReponse struct {
	ID        int64             `json:"id"`
	Email     string            `json:"email"`
	Name      string            `json:"name"`
	Realm     *GetRealmResponse `json:"realm"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
