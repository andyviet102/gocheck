package main

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func main() {
	app := fiber.New()

	// 1. Kết nối Redis (Dùng URL ông đưa)
	// Lưu ý: Nếu có mật khẩu, hãy điền vào, nếu không để trống ""
	rdb := redis.NewClient(&redis.Options{
		Addr: "shortline.proxy.rlwy.net:39547",
		Password: "", // Điền password nếu Railway yêu cầu
		DB:       0,  // DB mặc định
	})

	// Route mặc định (Root)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "UP", "redis": "connected"})
	})

	// 2. Chức năng ADD KEY (Method POST)
	// Body mẫu: {"key": "name", "value": "Viet"}
	app.Post("/set", func(c *fiber.Ctx) error {
		data := make(map[string]string)
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Lỗi parse JSON"})
		}

		err := rdb.Set(ctx, data["key"], data["value"], 0).Err()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"message": "Đã lưu thành công!"})
	})

	// 3. Chức năng READ KEY (Method GET)
	// URL mẫu: /get/name
	app.Get("/get/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		val, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy key"})
		} else if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"key": key, "value": val})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen("0.0.0.0:" + port)
}