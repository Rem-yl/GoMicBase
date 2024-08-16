package main

import (
	"Account/AccountWeb/handler"
	"Account/AccountWeb/middleware"
	conf "Account/Conf"
	logger "Account/Log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	config := conf.LoadConfig()
	host := config.GetString("web.host")
	port := config.GetString("web.port")
	dsn := fmt.Sprintf("%s:%s", host, port)

	r := gin.Default()

	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "test ok!",
			})
		})
		accountGroup.GET("/id/:id", handler.GetAccountByIdHandler)
		accountGroup.GET("/name/:name", handler.GetAccountByNameHandler)
		accountGroup.GET("/phone/:phone", handler.GetAccountByPhoneHandler)
		accountGroup.GET("/create", handler.CreateNewAccountHandler)
		accountGroup.POST("/login", handler.LoginHandler)
		accountGroup.GET("/jwt_test", middleware.JWTAuthMiddleware(), handler.JWTTestHandler)
	}

	r.Run(dsn)
}
