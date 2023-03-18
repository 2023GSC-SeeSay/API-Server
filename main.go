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

	err := app.Listen(":8080")
	//dd
	if err != nil {
		fmt.Printf("error starting server: %v", err)
		panic(err)
	}

}