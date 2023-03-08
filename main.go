package main

import (
	fiber "github.com/gofiber/fiber/v2"
)



func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:uid/audio", func(c *fiber.Ctx) error {
		uid := c.Params("uid")
		return c.SendString("Hello, " + uid + "!" +stt("C:\\workspace\\API-Server\\음성 002.wav", "C:\\workspace\\API-Server\\auth\\seesay-firebase-adminsdk-clpnw-faf918ab9f.json"))
	})

	app.Listen(":3000")

}