package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	// type H map[string]any
	c.JSON(http.StatusOK, gin.H{
		"rem": 20,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", hello)

	r.Run()
}
