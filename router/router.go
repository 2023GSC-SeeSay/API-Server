package router

import (
	handler "api-server/handler"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRouter (app *fiber.App) {
	api := app.Group("/api")

	api.Get("/audio/:uid/:pid", handler.AudioHandler)
}	