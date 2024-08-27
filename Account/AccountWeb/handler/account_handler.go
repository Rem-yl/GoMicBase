package handler

import (
	"Account/AccountServ/model"
	"Account/AccountServ/pb"
	"Account/AccountWeb/jwt_op"
	"Account/internal"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	share "github.com/GoMicBase/Share"

	"github.com/gin-gonic/gin"
)

func getAccountServClient(ctx *gin.Context) (client pb.AccountServiceClient) {
	client, err := internal.GetAccountServClient()
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

	client := getAccountServClient(ctx)
	if client == nil {
		return
	}
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

	client := getAccountServClient(ctx)
	if client == nil {
		return
	}
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

	client := getAccountServClient(ctx)
	if client == nil {
		return
	}
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

type newAccount struct {
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
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
	req := &pb.CreateAccountRequest{
		Name:     account.Name,
		Phone:    account.Phone,
		Password: account.Password,
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

func LoginHandler(ctx *gin.Context) {
	var account newAccount
	err := ctx.ShouldBind(&account)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  share.ErrParseAccount + err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	client := getAccountServClient(ctx)
	if client == nil {
		return
	}
	req := &pb.CheckNamePasswordRequest{
		Name:     account.Name,
		Password: account.Password,
	}

	resp, err := client.CheckNamePassword(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	if !resp.Check {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  share.ErrNotRegister,
			"data": "",
		})
		ctx.Abort()
		return
	}

	tokenStr, err := jwt_op.GenJWTToken(account.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  share.ErrGenJWTFailed + err.Error(),
			"data": "",
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"data": gin.H{
			"token": tokenStr,
		},
	})
}

func JWTTestHandler(ctx *gin.Context) {
	name := ctx.MustGet("name").(string)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "JWT ok",
		"data": gin.H{
			"name": name,
		},
	})
}

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
