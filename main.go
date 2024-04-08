package main

import (
	"fkclaude/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Forwarded-For",
	})
	router.SetupRoutes(app)
	app.Listen(":3650")
}
