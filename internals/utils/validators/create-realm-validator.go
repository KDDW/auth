package validators

import (
	"auth-service/internals/dtos"

	"github.com/go-playground/validator/v10"
)

func ValidateCreateRealmDto(dto *dtos.CreateRealmDto) []string {

	v := validator.New()
	err := v.Struct(dto)

	if err != nil {
		return []string{err.Error()}
	}
	return nil
}
