package router

import (
	handler "api-server/handler"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRouter (app *fiber.App) {
	api := app.Group("/api")

	// define routes
	api.Get("/audio/:uid/:pid", handler.AudioHandler)
	app.Get("/problems/:pid", handler.ProblemHandler) // fetching problem
	app.Post("/upload", handler.UploadGIFHandler) // uploading gif
}	


// curl -X POST -F "file=@C:\Users\HongEunbeen\Desktop\test.gif" http://localhost:8080/upload