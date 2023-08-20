package httpserver

import (
	"auth-service/internals/adapters/controllers"
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
)

func CreateServer() *fiber.App {

	app := fiber.New(fiber.Config{
		AppName:           "KDDW Auth Service",
		CaseSensitive:     false,
		EnablePrintRoutes: true,
	})

	SetupLoggerMiddleware(app)

	return app
}

func Listen(app *fiber.App) {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	address := ":" + port

	err := app.Listen(address)

	if err != nil {
		log.Fatal("cannot initialize fiber app: ", err)
	}
}

func RegisterControllers(app *fiber.App, controllers *controllers.Controllers) {
	registerRealmControllers(app, controllers)
	registerUsersControllers(app, controllers)
	registerAuthControllers(app, controllers)
}
