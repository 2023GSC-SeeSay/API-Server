package router

import fiber "github.com/gofiber/fiber/v2"

type CustomContext struct {
	*fiber.Ctx
	credFilePath string	
}