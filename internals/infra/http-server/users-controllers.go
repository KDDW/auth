package httpserver

import (
	"auth-service/internals/adapters/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func registerUsersControllers(app *fiber.App, controllers *controllers.Controllers) {

	router := app.Group("/api")

	router.Post("/users", controllers.UserController.CreateUser)
	router.Get("/users/:id", controllers.UserController.GetUserById)
	router.Delete("/users/:id", controllers.UserController.DeleteUser)
	router.Patch("/users/:id", controllers.UserController.UpdateUser)
	router.Get("/users", controllers.UserController.ListUsers)
}
