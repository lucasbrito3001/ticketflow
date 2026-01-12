package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbrito3001/go-kit/observability/httpctx"
	"github.com/lucasbrito3001/go-kit/observability/logger"
)

func main() {
	logger.Init("ticketflow-reservation-service")
	router := gin.Default()
	router.Use(httpctx.GinContextMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
