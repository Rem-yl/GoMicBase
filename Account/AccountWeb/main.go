package main

import (
	"Account/AccountWeb/handler"
	"Account/AccountWeb/middleware"
	logger "Account/Log"
	"Account/internal"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	accountWebConf := internal.AccountConf.AccountWebConf

	dsn := fmt.Sprintf("%s:%d", accountWebConf.Host, accountWebConf.Port)
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
