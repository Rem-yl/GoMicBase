package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `uri:"name" binding:"required"` // https://gin-gonic.com/docs/examples/binding-and-validation/
	Age  int    `uri:"age" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:age", indexHandle)

	r.Run(":8081")
}

func indexHandle(c *gin.Context) {
	var p Person
	if err := c.ShouldBindUri(&p); err != nil { // 注意这里使用的方法是 ShouldBindUri()
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get",
		"name": p.Name,
		"age":  p.Age,
	})
}
