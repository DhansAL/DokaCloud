package router

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	// Test Routes
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})
}
