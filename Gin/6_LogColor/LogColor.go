package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ping")
	})

	r.Run()
}
