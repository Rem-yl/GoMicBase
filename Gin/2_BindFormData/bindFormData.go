package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"` // 这里Name和Age如果不是大写开头就无法正确绑定
	Age  string `form:"age"`
}

func RemFunc(ctx *gin.Context) {
	var p Person
	ctx.Bind(&p)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "rem",
		"name": p.Name,
		"age":  p.Age,
	})
}

func RamFunc(ctx *gin.Context) {
	var p Person
	ctx.Bind(&p)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ram",
		"name": p.Name,
		"age":  p.Age,
	})
}

// curl localhost:8081/rem?name=rem&age=10
func main() {
	r := gin.Default()
	r.GET("/rem", RemFunc)
	r.GET("/ram", RamFunc)
	r.Run(":8081")
}
