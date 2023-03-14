package main

import (
	router "api-server/router"
	"fmt"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {

	app := fiber.New()

	// load cred file path from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	gcp_cred_file_path := os.Getenv("GCP_CRED_FILE_PATH")
	fmt.Println(gcp_cred_file_path)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("<h1>SeeSay API Server</h1>")
	// })

	// app.Use(middleware.Logger())

	router.SetupRouter(app, gcp_cred_file_path)

	app.Listen(":3000")

}