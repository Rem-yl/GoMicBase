package handler

import (
	"Account/proto/pb"
	"AccountWeb/req"
	"AccountWeb/res"
	"context"
	"crypto/md5"
	"fmt"
	"jwt_op"
	"net/http"
	"strconv"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
)

var DefaultOptions = password.Options{
	SaltLen:      16,
	Iterations:   100,
	KeyLen:       32,
	HashFunction: md5.New,
}

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

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Id       uint32 `json:"id"`
}

func AccountLogin(ctx *gin.Context) {
	var user UserLogin
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Conn err :%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)
	req := &pb.CheckAccountRequest{
		Nickname: user.Name,
		Password: user.Password,
		Id:       user.Id,
	}

	resp, err := client.CheckNamePassword(context.Background(), req)
	if err != nil {
		fmt.Printf("CheckNamePassword Failed: %s\n", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"Check": resp.Check,
	})
}

func LoginByPassword(ctx *gin.Context) {
	var loginByPasword req.LoginByPassword
	err := ctx.ShouldBind(&loginByPasword)
	if err != nil {
		fmt.Println("LoginByPassword解析参数出错: " + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "参数解析错误",
		})
		return
	}

	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		fmt.Println("LoginByPassword拨号出错: " + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "拨号错误",
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)
	resp, err := client.GetAccountByPhoneNumber(context.Background(), &pb.PhoneNumberRequest{
		PhoneNumber: loginByPasword.PhoneNumber,
	})
	if err != nil {
		fmt.Println("GetAccountByPhoneNumber出错: " + err.Error())
		return
	}

	fmt.Printf("xxxxxx: resp: %v\n", resp)
	checkResp, err := client.CheckNamePassword(context.Background(), &pb.CheckAccountRequest{
		Nickname: resp.Nickname,
		Password: resp.Password,
		Id:       resp.Id,
	})
	fmt.Printf("checkResp: %v\n", checkResp)
	if err != nil {
		fmt.Println("CheckNamePassword出错: " + err.Error())
		return
	}
	checkResult := "登录失败"
	if checkResp.Check {
		checkResult = "登录成功"
		j := jwt_op.NewJWT()
		now := time.Now()
		claims := jwt_op.CustomClaims{
			StandardClaims: jwt.StandardClaims{
				NotBefore: now.Unix(),
				ExpiresAt: now.Add(time.Hour * 24 * 30).Unix(),
			},
			Id:          resp.Id,
			Nickname:    resp.Nickname,
			AuthorityId: int32(resp.Role),
		}
		token, err := j.GenerateJWT(claims)
		if err != nil {
			fmt.Println("GenerateJWT 出错: " + err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":    "",
			"result": checkResult,
			"token":  token,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "",
		"result": checkResult,
		"token":  "",
	})
}
