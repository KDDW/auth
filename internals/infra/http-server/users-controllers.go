package httpserver

import (
	"auth-service/internals/adapters/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func registerUsersControllers(app *fiber.App, controllers *controllers.Controllers) {

	app.Post("/users", controllers.UserController.CreateUser)
	app.Get("/users/:id", controllers.UserController.GetUserById)

}
