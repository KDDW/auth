package dtos

type CreateRealmDto struct {
	Code string `json:"code" validate:"required"`
}
