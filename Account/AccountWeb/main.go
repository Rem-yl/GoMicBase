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

	share "github.com/GoMicBase/Share"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	accountWebConf := conf.AccountConf.AccountWebConf

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

	r.GET("/health", handler.HealthHandler)

	err := internal.ConsulRegWeb(accountWebConf.Host, int(accountWebConf.Port), accountWebConf.Name, accountWebConf.Id, []string{"test"})
	if err != nil {
		log.Panicf("%s:%s\n", share.ErrWebRegister, err.Error())
	}
	r.Run(dsn)
}
