package middleware

import (
	"jwt_op"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" || len(token) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "认证失败, 需要登录",
			})
			ctx.Abort()
			return
		}

		j := jwt_op.NewJWT()
		parserToken, err := j.ParseToken(token)
		if err != nil {
			if err.Error() == jwt_op.ErrTokenExpired.Error() {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"msg": jwt_op.ErrTokenExpired,
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "认证失败, 需要登录",
			})
			ctx.Abort()
			return
		}
		ctx.Set("claims", parserToken)
		ctx.Next()

	}
}
