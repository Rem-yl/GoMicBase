package middleware

import (
	"Account/AccountWeb/jwt_op"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Auth")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "empty jwt auth",
			})
			ctx.Abort()
			return
		}

		mc, err := jwt_op.ParseJWTToken(authHeader)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "无效token",
				"data": "{}",
			})
			ctx.Abort()
			return
		}

		ctx.Set("name", mc.Name)
		ctx.Next()
	}
}
