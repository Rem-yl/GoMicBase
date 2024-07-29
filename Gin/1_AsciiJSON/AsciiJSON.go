package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl localhost:9090/asciijson
func main() {
	r := gin.Default()
	r.GET("/asciijson", func(ctx *gin.Context) {
		data := map[string]string{
			"name": "小鱼干", // ctx.AsciiJSON方法只输出ascii字符
			"tag":  "</br>",
		}

		// Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
		ctx.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
