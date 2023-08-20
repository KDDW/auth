package validators

import (
	"auth-service/internals/dtos"
)

func ValidateCreateRealmDto(dto *dtos.CreateRealmDto) []string {
	return getErrorsFromValidation(dto)
}
