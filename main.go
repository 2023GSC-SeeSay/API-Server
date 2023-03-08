package main

import (
	fiber "github.com/gofiber/fiber/v2"
)



func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("<h1>SeeSay API Server</h1>")
	})
	app.Route("/api", func(r fiber.Router) {
		r.Use("/audio", stt)
	})

	app.Listen(":3000")

}