package httpserver

import (
	"auth-service/internals/adapters/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func registerRealmControllers(app *fiber.App, controllers *controllers.Controllers) {

	router := app.Group("/api")

	router.Post("/realms", controllers.RealmController.CreateRealm)
	router.Get("/realms", controllers.RealmController.ListRealms)
	router.Get("/realms/:id", controllers.RealmController.GetRealmById)
}
