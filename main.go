package main

import (
	"github.com/gofiber/fiber/v2"
)

type JSONTextResponse struct {
	Message string
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(JSONTextResponse{Message: "hello fiber"})

	})

	app.Listen(":8080")
}
