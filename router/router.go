package router

import (
	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	//* api test hit
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "success hit api",
		})
	})

	app.Post("/login", handlers.Login)

	protected := app.Group("/protected")
	protected.Get("/data", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "this is protected data",
		})
	})
}
