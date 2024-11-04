package main

import (
	"github.com/agill7227/Web-Application-Firewall/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./Views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(middleware.WAFMiddleware)

	app.Static("/", "./Views")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	app.Post("/data", func(c *fiber.Ctx) error {
		return c.SendString("Data received!")
	})

	app.Listen(":3000")

}
