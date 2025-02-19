package router

import (
	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/handlers"
	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/middlewares"
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

	protectedRoute := app.Group("/protected")
	protectedRoute.Get("get-logged-user", middlewares.IsAuth(), handlers.GetLoggedUser)
}
