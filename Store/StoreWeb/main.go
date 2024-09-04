package main

import (
	conf "Store/Conf"
	"Store/StoreWeb/handler"
	"fmt"
	"log"
	"net/http"

	share "github.com/GoMicBase/Share"
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

	consulClient, err := share.GetConsulClient(consulConf)
	if err != nil {
		log.Panicln(err.Error())
	}

	err = share.ConsulRegWeb(consulClient, storeWebConf.Host, int(storeWebConf.Port), storeWebConf.Name, storeWebConf.Id, []string{"test"})
	if err != nil {
		log.Panicln(err.Error())
	}

	r.Run(dsn)
}
