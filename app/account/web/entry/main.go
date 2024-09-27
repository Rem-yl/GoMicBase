package main

import (
	"GoMicBase/app/account/web/handler"
	"GoMicBase/pkg/registry"
	"GoMicBase/pkg/zlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "test ok!",
			})
		})
		accountGroup.GET("/create", handler.CreateNewAccountHandler)
	}

	r.GET("/health", handler.HealthHandler)

	register, _ := registry.NewConsulRegistery("127.0.0.1", 8500)

	register.RegisterWeb("127.0.0.1", 8080, "accountWeb", "rem123")
	dsn := "127.0.0.1:8080"
	zlog.Infof("Start Web on: %s", dsn)
	go func() {
		if err := r.Run(dsn); err != nil {
			log.Panicln(err.Error())
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	log.Println("Received interrupt signal, shutting down gracefully...")

	if err := register.Deregister("rem123"); err != nil {
		log.Println(err.Error())
	}

	log.Println("Account Web shutdown.")
}
