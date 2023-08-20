package validators

import (
	"auth-service/internals/dtos"
)

func ValidateCreateUserDto(dto *dtos.CreateUserDto) []string {
	return getErrorsFromValidation(dto)
}
