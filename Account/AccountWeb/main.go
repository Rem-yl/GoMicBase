package main

import (
	"Account/AccountWeb/handler"
	"Account/AccountWeb/middleware"
	conf "Account/Conf"
	logger "Account/Log"
	"Account/internal"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	var nacosConfig conf.NacosConfig
	var accountWebConfig conf.AccountWebConfig
	if err := internal.LoadAccountWebConfig("./conf", "dev", &nacosConfig, &accountWebConfig); err != nil {
		log.Panicln(err.Error())
	}
	dsn := fmt.Sprintf("%s:%d", accountWebConfig.Host, accountWebConfig.Port)
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
