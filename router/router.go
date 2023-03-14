package router

import (
	handler "api-server/handler"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRouter (app *fiber.App, cred_file_path string) {
	api := app.Group("/api")

	api.Post("/audio/:uid/:pid", handler.AudioHandler(cred_file_path))
}	