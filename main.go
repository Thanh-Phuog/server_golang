package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/list-product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List Product",
		})
	})

	r.GET("/api/list-user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List User",
		})
	})

	r.GET("/api/list-user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "List User",
			"id":      id,
		})
	})
	r.Run(":3000")
}
