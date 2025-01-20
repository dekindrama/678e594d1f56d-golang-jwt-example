package main

import (
	"fmt"

	"github.com/dekindrama/678e594d1f56d-golang-jwt-example/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)

	//* run fiber instance
	fmt.Println("server run on http://localhost:8000")
	app.Listen(":8000")
}
