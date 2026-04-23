package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Đổi từ app.Get("/health", ...) thành app.Get("/", ...)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "UP",
			"message": "tao là bố của chúng mày",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}