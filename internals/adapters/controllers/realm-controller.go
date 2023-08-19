package controllers

import (
	realm_services "auth-service/internals/domain/services/realm-services"
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"auth-service/internals/utils/validators"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RealmController struct {
	services *realm_services.RealmServices
}

func NewRealmController(s *realm_services.RealmServices) *RealmController {
	return &RealmController{
		services: s,
	}
}

func (c *RealmController) CreateRealm(ctx *fiber.Ctx) error {

	dto := &dtos.CreateRealmDto{}

	err := ctx.BodyParser(&dto)

	var authError *exceptions.AuthError = nil

	if err != nil {
		msg := "can't parse request: " + err.Error()
		authError = exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	validationErrors := validators.ValidateCreateRealmDto(dto)

	if validationErrors != nil {
		msg := strings.Join(validationErrors, ", ")
		authError = exceptions.NewAuthError(400, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	authError = c.services.CreateRealm(dto)

	if authError != nil {
		return exceptions.AuthErrorToFiberError(authError)
	}

	return ctx.SendStatus(201)
}

func (c *RealmController) ListRealms(ctx *fiber.Ctx) error {

	params := &dtos.ListRealmsDto{}
	err := ctx.QueryParser(params)

	if err != nil {
		msg := "cannot parse query parameters: " + err.Error()
		authError := exceptions.NewAuthError(500, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}

	realms, authError := c.services.ListRealms(params)

	if authError != nil {
		return exceptions.AuthErrorToFiberError(authError)
	}

	out, err := json.Marshal(realms)

	if err != nil {
		msg := "cannot json.Marshal on ListRealms controller: " + err.Error()
		authError := exceptions.NewAuthError(500, msg)
		return exceptions.AuthErrorToFiberError(authError)
	}
	return ctx.Send(out)
}
