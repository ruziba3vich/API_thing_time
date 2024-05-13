package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/time_project/backend/internal/handlers"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.POST("/time/RFC3339", func(c *gin.Context) {
		handlers.GetResult(c)
	})

	router.Run(":7777")
}
