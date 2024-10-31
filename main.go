package main

import (
	"github.com/agill7227/Web-Application-Firewall/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(middleware.WAFMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Test test")
	})
	app.Post("/data", func(c *fiber.Ctx) error {
		return c.SendString("Data received!")
	})

	app.Listen(":3000")

}
