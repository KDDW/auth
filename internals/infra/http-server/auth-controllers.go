package httpserver

import (
	"auth-service/internals/adapters/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func registerAuthControllers(app *fiber.App, controllers *controllers.Controllers) {

	app.Post("/auth/login", controllers.AuthController.Login)
	app.Post("/auth/login/refresh", controllers.AuthController.LoginRefreshToken)
	app.Post("/auth/token/verify", controllers.AuthController.VerifyToken)
}
