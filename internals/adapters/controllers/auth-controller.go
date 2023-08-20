package controllers

import (
	auth_services "auth-service/internals/domain/services/auth-services"
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"auth-service/internals/utils/validators"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	services *auth_services.AuthServices
}

func NewAuthController(services *auth_services.AuthServices) *AuthController {
	return &AuthController{
		services: services,
	}
}

func (a *AuthController) Login(c *fiber.Ctx) error {

	dto := &dtos.LoginRequestDto{}

	err := c.BodyParser(&dto)

	if err != nil {
		msg := "can't parse request: " + err.Error()
		authError := exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	validationErrors := validators.ValidateLoginRequest(dto)

	if validationErrors != nil {
		msg := strings.Join(validationErrors, ", ")
		authError := exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	loginResponse, authError := a.services.Login(dto.Email, dto.Password, dto.Realm)

	if authError != nil {
		return exceptions.AuthErrorToFiberError(authError)
	}

	jsonOut, err := json.Marshal(loginResponse)

	if err != nil {
		fmt.Print("error while json.Marshal in login request: ", err.Error())
		authError := exceptions.NewAuthError(500, "unknow error")
		return exceptions.AuthErrorToFiberError(authError)
	}

	c.Status(200).Send(jsonOut)

	return nil
}

func (a *AuthController) VerifyToken(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()

	authorization := headers["Authorization"]

	if authorization == "" {
		authError := exceptions.NewAuthError(400, "missing bearer token")
		return exceptions.AuthErrorToFiberError(authError)
	}

	tokenData := strings.Split(authorization, " ")

	if len(tokenData) != 2 {
		authError := exceptions.NewAuthError(400, "invalid bearer format")
		return exceptions.AuthErrorToFiberError(authError)
	}

	if tokenData[0] != "Bearer" {
		authError := exceptions.NewAuthError(400, "invalid bearer format")
		return exceptions.AuthErrorToFiberError(authError)
	}

	token := tokenData[1]

	isValid, authErr := a.services.VerifyToken(token)

	if authErr != nil {
		return exceptions.AuthErrorToFiberError(authErr)
	}

	if isValid {
		c.SendStatus(200)
	} else {
		authErr := exceptions.NewAuthError(500, "internal error")
		return exceptions.AuthErrorToFiberError(authErr)
	}

	return nil
}

func (a *AuthController) LoginRefreshToken(c *fiber.Ctx) error {

	dto := &dtos.LoginRefreshTokenDto{}

	err := c.BodyParser(&dto)

	if err != nil {
		msg := "can't parse request: " + err.Error()
		authError := exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	validationErrors := validators.ValidateLoginRefreshTokenDto(dto)

	if validationErrors != nil {
		msg := strings.Join(validationErrors, ", ")
		authError := exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	res, authErr := a.services.LoginRefreshToken(dto.RefreshToken)

	if authErr != nil {
		return exceptions.AuthErrorToFiberError(authErr)
	}

	out, err := json.Marshal(res)

	if err != nil {
		msg := "error while json.Marshal in LoginRefreshToken: " + err.Error()
		authError := exceptions.NewAuthError(500, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	c.Status(200).Send(out)

	return nil
}
