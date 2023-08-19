package controllers

import (
	user_services "auth-service/internals/domain/services/user-services"
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"auth-service/internals/utils/mappers"
	"auth-service/internals/utils/validators"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	services *user_services.UserServices
}

func NewUserController(s *user_services.UserServices) *UserController {
	return &UserController{
		services: s,
	}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {

	dto := &dtos.CreateUserDto{}

	err := ctx.BodyParser(&dto)

	if err != nil {
		msg := "can't parse request: " + err.Error()
		authError := exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	validationErrors := validators.ValidateCreateUserDto(dto)

	if validationErrors != nil {
		msg := strings.Join(validationErrors, ", ")
		authError := exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	authError := c.services.CreateUser(dto)

	if authError != nil {
		return exceptions.AuthErrorToFiberError(authError)
	}

	return nil
}

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		msg := "can't get user without and id"
		authError := exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)

	}

	idNbr, err := strconv.ParseUint(id, 0, 64)

	if err != nil {
		authError := exceptions.NewAuthError(400, "invalid id")
		return exceptions.AuthErrorToFiberError(authError)
	}

	user, authError := c.services.GetById(int64(idNbr))

	if authError != nil {
		return exceptions.AuthErrorToFiberError(authError)
	}

	res := mappers.MapUserModelToResponse(user)

	out, err := json.Marshal(res)

	if err != nil {
		msg := "cannot json.Marshal on GetUserById controller: " + err.Error()
		authError := exceptions.NewAuthError(500, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	return ctx.Status(200).Send(out)
}
