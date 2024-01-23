package main

import (
	"github.com/cancodes/fiber-templ-htmx-tailwindcss-template/components"
	"github.com/cancodes/fiber-templ-htmx-tailwindcss-template/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.Render(c, components.Hello(""))
	})
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return handlers.Render(c, components.Name(name))
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
