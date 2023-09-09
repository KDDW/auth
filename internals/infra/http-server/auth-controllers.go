package httpserver

import (
	"auth-service/internals/adapters/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func registerAuthControllers(app *fiber.App, controllers *controllers.Controllers) {

	router := app.Group("/api")

	router.Post("/auth/login", controllers.AuthController.Login)
	router.Post("/auth/login/refresh", controllers.AuthController.LoginRefreshToken)
	router.Post("/auth/token/verify", controllers.AuthController.VerifyToken)
}
