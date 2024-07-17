package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	productGroup := r.Group("/product")
	{
		productGroup.GET("/productList", productList)
		productGroup.GET("/:id", productIdDetail)
	}

	r.Run()
}

func productList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"rem":  10,
		"ram":  20,
		"fish": 30,
	})
}

func productIdDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad id",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}
