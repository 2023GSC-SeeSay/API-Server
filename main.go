package main

import (
	fiber "github.com/gofiber/fiber/v2"

	router "api-server/router"
)



func main() {

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("<h1>SeeSay API Server</h1>")
	// })

	// app.Use(middleware.Logger())

	router.SetupRouter(app)

	app.Listen(":3000")

}