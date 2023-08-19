package httpserver

import (
	"auth-service/internals/utils/terminal"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupLoggerMiddleware(app *fiber.App) {

	format := terminal.Colors["cyan"] + "${time}" + terminal.Colors["reset"]
	format += " [${status}] ${method} ${path}?${queryParams}"
	format += "\t${resBody}\n"

	app.Use(logger.New(logger.Config{
		Format:     format,
		TimeZone:   "UTC",
		TimeFormat: "2006-01-02 15:04:05",
	}))
}
