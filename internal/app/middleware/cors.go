package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CORSMiddlewareは、CORSの設定を行うミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigin, ok := os.LookupEnv("CORS_ALLOWED_ORIGIN")
		if !ok {
			allowedOrigin = "http://localhost:3000"
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin) // ここを環境変数から取得
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Cookie や セッションを送信する場合は必須
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
