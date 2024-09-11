package main

import (
	conf "Store/Conf"
	"Store/StoreWeb/handler"
	"fmt"
	"log"
	"net/http"

	register "github.com/GoMicBase/Register"
	"github.com/gin-gonic/gin"
)

func main() {
	storeWebConf := conf.StoreConf.StoreWebConf
	consulConf := conf.StoreConf.ConsulConf
	dsn := fmt.Sprintf("%s:%d", storeWebConf.Host, storeWebConf.Port)
	r := gin.Default()

	brandGroup := r.Group("/brand")
	{
		brandGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "brand test ok!",
				"data": "{}",
			})
		})
	}
	r.GET("/health", handler.HealthHandler)
	consulRegistery := &register.ConsulRegistery{
		Config: &register.ConsulConfig{
			Host: consulConf.Host,
			Port: consulConf.Port,
		},
	}

	err := consulRegistery.NewClient()
	if err != nil {
		log.Println(err.Error())
	}

	err = consulRegistery.RegisterWeb(storeWebConf.Host, int(storeWebConf.Port), storeWebConf.Name, storeWebConf.Id, []string{"test"})
	if err != nil {
		log.Panicln(err.Error())
	}

	r.Run(dsn)
}
