package middleware

import (
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

		// parts := strings.SplitN(authHeader, " ", 2)

	}
}
