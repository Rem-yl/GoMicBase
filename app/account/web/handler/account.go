package handler

import (
	v1 "GoMicBase/api/account/service/v1"
	"GoMicBase/app/account/conf"
	"GoMicBase/app/account/model"
	"GoMicBase/app/account/server"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

type newAccount struct {
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func getAccountServClient(ctx *gin.Context) v1.AccountServiceClient {
	config, _ := conf.NewAccountConfig("../", "dev")
	client, err := server.GetGrpcClient(config.ConsulConfig, config.AccountServConfig)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return nil
	}

	return client
}

func CreateNewAccountHandler(ctx *gin.Context) {
	var account newAccount
	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	client := getAccountServClient(ctx)
	if client == nil {
		return
	}
	req := &v1.CreateAccountRequest{
		Name:     "ram",
		Phone:    "18801117212",
		Password: "1234567890",
	}
	resp, err := client.CreateAccount(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	accountResp := model.PbResp2CustomAccount(resp)
	jsonData, err := json.Marshal(accountResp)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	accountJSON := string(jsonData)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": accountJSON,
	})
}
