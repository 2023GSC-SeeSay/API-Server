package main

import (
	fiber "github.com/gofiber/fiber/v2"
)



func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("<h1>SeeSay API Server</h1>")
	})

	api := app.Group("/api")

	api.Get("/audio", handler.stt_handler)


	app.Listen(":3000")

}