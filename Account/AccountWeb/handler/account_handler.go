package handler

import (
	"Account/AccountServ/model"
	"Account/AccountServ/pb"
	conf "Account/Conf"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func getGrpcAddr() string {
	config := conf.LoadConfig()
	host := config.GetString("grpc.host")
	port := config.GetString("grpc.port")

	addr := fmt.Sprintf("%s:%s", host, port)
	return addr
}

func GetAccountByIdHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "empty id",
			"data": "",
		})
		ctx.Abort()
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	addr := getGrpcAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	client := pb.NewAccountServiceClient(conn)
	req := &pb.AccountIdRequest{
		Id: uint32(id),
	}
	resp, err := client.GetAccountById(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
	}

	account := model.PbResp2CustomAccount(resp)
	jsonData, err := json.Marshal(account)
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

func GetAccountByNameHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "empty name",
			"data": "",
		})
		ctx.Abort()
		return
	}

	addr := getGrpcAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	client := pb.NewAccountServiceClient(conn)
	req := &pb.AccountNameRequest{
		Name: name,
	}
	resp, err := client.GetAccountByName(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
	}

	account := model.PbResp2CustomAccount(resp)
	jsonData, err := json.Marshal(account)
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

func GetAccountByPhoneHandler(ctx *gin.Context) {
	phone := ctx.Param("phone")
	if phone == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "empty phone",
			"data": "",
		})
		ctx.Abort()
		return
	}

	addr := getGrpcAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	client := pb.NewAccountServiceClient(conn)
	req := &pb.AccountPhoneRequest{
		Phone: phone,
	}
	resp, err := client.GetAccountByPhone(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
	}

	account := model.PbResp2CustomAccount(resp)
	jsonData, err := json.Marshal(account)
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
