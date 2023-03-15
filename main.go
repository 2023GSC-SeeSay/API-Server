package main

import (
	router "api-server/router"

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

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}

}