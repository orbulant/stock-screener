package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Setup routes for the application

// SetupRoutes sets up all the routes for the application

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "OK",
		})
	})
}
