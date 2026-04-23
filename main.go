package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
    app := fiber.New()

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "UP",
        })
    })

   
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000" // Default port
    }

    app.Listen("0.0.0.0:" + port)
}