package httpserver

import (
	"auth-service/internals/adapters/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func registerRealmControllers(app *fiber.App, controllers *controllers.Controllers) {

	app.Post("/realms", controllers.RealmController.CreateRealm)
	app.Get("/realms", controllers.RealmController.ListRealms)
	app.Get("/realms/:id", controllers.RealmController.GetRealmById)
}
