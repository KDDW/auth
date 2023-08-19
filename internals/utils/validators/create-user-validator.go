package validators

import (
	"auth-service/internals/dtos"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateCreateUserDto(dto *dtos.CreateUserDto) []string {

	var errors []string

	v := validator.New()
	err := v.Struct(dto)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			errorStr := fmt.Sprintf("%s failed on validation of '%s'", err.Field(), err.Tag())
			errors = append(errors, errorStr)
		}
		return errors
	}
	return nil
}
