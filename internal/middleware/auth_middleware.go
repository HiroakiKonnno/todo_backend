package middleware

import (
	"net/http"
	auth "todo_backend/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("jwt")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "トークンが見つかりません"})
			ctx.Abort()
			return
		}
		_, err = auth.ValidateJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "無効なトークン"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
