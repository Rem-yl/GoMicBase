package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name") // 获取url中的name参数
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello, " + name,
		})
	})

	r.Run()
}
