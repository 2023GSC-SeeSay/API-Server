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
	app.Post("/bookshelf", handler.BookshelfHandler) // uploading bookshelf
	app.Post("/upload", handler.UploadGIFHandler) // uploading gif

}	

// Test upolad gif
// curl -X POST -F "file=@C:\Users\HongEunbeen\Desktop\test.gif" http://localhost:8080/upload

// Test bookshelf
// curl -X POST -H "Content-Type: application/json" -d "{\"pid\": 1, \"uid\": 1, \"text\": \"hello world!\"}" http://localhost:8080/bookshelf

