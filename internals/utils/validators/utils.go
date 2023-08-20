package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func getErrorsFromValidation(dto interface{}) []string {

	var errors []string

	if validate == nil {
		validate = validator.New()
	}

	err := validate.Struct(dto)

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
