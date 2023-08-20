package exceptions

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type AuthError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Details string `json:"details"`
}

func getMessageFromCode(code int) string {

	switch code {
	case 409:
		return "Conflict exception"
	case 400:
		return "Bad request exception"
	case 401:
		return "Unauthorized exception"
	case 404:
		return "Not found exception"
	case 500:
		return "Internal server exception"
	default:
		return "Unknow exception"
	}

}

func NewAuthError(code int, message string) *AuthError {

	return &AuthError{
		Message: getMessageFromCode(code),
		Code:    code,
		Details: message,
	}

}

func AuthErrorToFiberError(authError *AuthError) *fiber.Error {

	authErrorBytes, err := json.Marshal(authError)

	if err != nil {
		msg := "error while json.marshal auth error: " + err.Error()
		return fiber.NewError(500, msg)
	}

	return fiber.NewError(authError.Code, string(authErrorBytes))
}
