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



	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis.railway.internal:6379", 
		Username: "default",                        
		Password: "dlhxpPzrdgBcLDsxqGzEhozdzEMlytiL", 
		DB:       0,
	})
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "UP",
			"message": "tao là bố của chúng mày",
		})
	})

	app.Post("/set", func(c *fiber.Ctx) error {
		data := make(map[string]string)
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "JSON không hợp lệ"})
		}

		err := rdb.Set(ctx, data["key"], data["value"], 0).Err()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Lỗi lưu Redis: " + err.Error()})
		}

		return c.JSON(fiber.Map{"message": "Đã lưu " + data["key"] + " thành công!"})
	})

	
	app.Get("/get/:key", func(c *fiber.Ctx) error {
		key := c.Params("key")
		val, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			return c.Status(404).JSON(fiber.Map{"error": "Key không tồn tại"})
		} else if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Lỗi đọc Redis"})
		}

		return c.JSON(fiber.Map{"key": key, "value": val})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}