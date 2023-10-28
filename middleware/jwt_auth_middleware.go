package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/washington-shoji/gin-api/helpers"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helpers.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Access denied")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func JwtAuthMiddlewareCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helpers.TokenValidCookie(ctx)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/api/auth")
			return
		}
		ctx.Next()
	}
}
