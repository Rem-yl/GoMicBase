package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl localhost:9091/ping
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ping",
		})
	})

	r.Run("localhost:9091")
}
