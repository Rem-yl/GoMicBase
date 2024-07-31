package handler

import (
	"Account/proto/pb"
	"AccountWeb/res"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func AccountListHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Connect Error :%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)
	resp, err := client.GetAccountList(context.Background(), &pb.PageinRequest{
		PageNumber: 1,
		PageSize:   3,
	})
	if err != nil {
		fmt.Printf("GetAccountList Failed: %s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	var resAccountList []res.AccountRes
	for _, accountResp := range resp.AccountList {
		resAccount := res.Pb2model(accountResp)
		resAccountList = append(resAccountList, resAccount)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"total": resp.Total,
		"data":  resAccountList,
	})
}

func GetAccountIdHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Connect Error :%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("GetAccountId Parse Id: %s Failed", idStr)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)
	resp, err := client.GetAccountById(context.Background(), &pb.IdRequest{
		Id: uint32(id),
	})

	if err != nil {
		fmt.Printf("GetAccountId Failed: %s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resAccount := res.Pb2model(resp)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": resAccount,
	})
}
