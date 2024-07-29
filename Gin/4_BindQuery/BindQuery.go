package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/query", queryHandle)
	r.POST("/post", postHandle)

	r.Run(":8081")
}

func postHandle(c *gin.Context) {
	var p Person
	err := c.ShouldBind(&p)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "post",
		"name": p.Name,
		"age":  p.Age,
	})
}

func queryHandle(c *gin.Context) {
	var p Person
	err := c.ShouldBind(&p)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "GetQuery",
		"name": p.Name,
		"age":  p.Age,
	})

}
