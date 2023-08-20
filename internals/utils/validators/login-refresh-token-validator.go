package validators

import (
	"auth-service/internals/dtos"
)

func ValidateLoginRefreshTokenDto(dto *dtos.LoginRefreshTokenDto) []string {
	return getErrorsFromValidation(dto)
}
