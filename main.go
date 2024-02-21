package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	 "github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./web/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./web/static")

	app.Get("/", func(c *fiber.Ctx) error {
		engine.Reload(true) // For dev use only. Remove this line in production.
		return c.Render("index", fiber.Map{})
	})

	log.Fatal(app.Listen("127.0.0.1:5335"), "Failed to start server")
}
