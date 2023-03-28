package main

import (
	router "github.com/2023GSC-SeeSay/API-Server/router"

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
	//dd
	if err != nil {
		panic(err)
	}

}
