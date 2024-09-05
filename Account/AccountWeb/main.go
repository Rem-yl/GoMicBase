package main

import (
	"Account/AccountWeb/handler"
	"Account/AccountWeb/middleware"
	conf "Account/Conf"
	logger "Account/Log"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	register "github.com/GoMicBase/Register"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	accountWebConf := conf.AccountConf.AccountWebConf
	consulConf := conf.AccountConf.ConsulConf

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

	// consulClient, err := share.GetConsulClient(share.ConsulConfig(consulConf))
	consulRegistery := &register.ConsulRegistery{
		Config: &register.ConsulConfig{
			Host: consulConf.Host,
			Port: consulConf.Port,
		},
	}

	err := consulRegistery.NewClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	if err = consulRegistery.RegisterWeb(accountWebConf.Host, int(accountWebConf.Port), accountWebConf.Name, accountWebConf.Id, []string{"test"}); err != nil {
		log.Panicf(err.Error())
	}
	log.Printf("Start Account Web on: %s:%d", accountWebConf.Host, accountWebConf.Port)

	go func() {
		if err := r.Run(dsn); err != nil {
			log.Panicln(err.Error())
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	log.Println("Received interrupt signal, shutting down gracefully...")

	if err = consulRegistery.Deregister(accountWebConf.Id); err != nil {
		log.Println(err.Error())
	}

	log.Println("Account Web shutdown.")

}
