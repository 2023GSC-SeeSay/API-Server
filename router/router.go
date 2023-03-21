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
	app.Get("/gif/:gif_path", handler.GIFHandler) // fetching gif
}	


// 아래 주소에서 동작 확인
// /problems/0
// /problems/1
// /gif/gif/4JU.gif
// /gif/gif/7IsD.gif
