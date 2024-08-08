package main

import (
	"AccountWeb/handler"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "ip addr")
	port := flag.Int("port", 8081, "port")
	addr := fmt.Sprintf("%s:%d", *ip, *port)

	r := gin.Default()
	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/list", handler.AccountListHandler)
		accountGroup.GET("/user/:id", handler.GetAccountIdHandler)
		accountGroup.POST("/login", handler.AccountLogin) // use phoneNumber & password
		accountGroup.POST("/jwtlogin", handler.LoginByPassword)
	}

	r.Run(addr)
}
