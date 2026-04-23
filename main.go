package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()

	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "UP",
			"timestamp": time.Now().Format(time.RFC3339),
			"service":   "my-awesome-service",
		})
	})


	r.Run(":8080")
}