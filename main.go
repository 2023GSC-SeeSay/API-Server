package main

import (
	"api-server/handler"
	router "api-server/router"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

// define custom context for google cloud credential file path
type CustomContext struct {
	*fiber.Ctx
	credFilePath string	
}

func main() {
	
	app := fiber.New()
	
	router.SetupRouter(app)
	app.Get("/problems/:pid", handler.ProblemHandler) // fetching problem
	app.Get("/gif/:gif_path", handler.GIFHandler) // fetching gif
	err := app.Listen(":8080")
	//dd
	if err != nil {
		fmt.Printf("error starting server: %v", err)
		panic(err)
	}

}

// 아래 주소에서 동작 확인
// /problems/0
// /problems/1
// /gif/gif/4JU.gif
// /gif/gif/7IsD.gif
