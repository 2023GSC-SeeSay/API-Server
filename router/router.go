package router

import (
	handler "api-server/handler"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRouter (app *fiber.App) {
	api := app.Group("/api")

	// define routes
	api.Get("/audio/:pid", handler.AudioHandler)
	api.Get("/problem/:pid", handler.ProblemHandler)
}	