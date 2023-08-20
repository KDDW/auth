package validators

import (
	"auth-service/internals/dtos"
)

func ValidateLoginRequest(dto *dtos.LoginRequestDto) []string {
	return getErrorsFromValidation(dto)
}
