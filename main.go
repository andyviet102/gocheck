package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Khởi tạo app Fiber
	app := fiber.New()

	// Route healthcheck
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "UP",
			"message": "tao là bố của chúng mày",
		})
	})

	// Lấy PORT từ biến môi trường của Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Mặc định chạy port 3000 nếu ở local
	}

	// QUAN TRỌNG: Phải lắng nghe trên 0.0.0.0 để Railway nhận được traffic
	app.Listen("0.0.0.0:" + port)
}