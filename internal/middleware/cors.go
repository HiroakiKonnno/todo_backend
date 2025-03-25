package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSMiddlewareは、CORSの設定を行うミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigin, ok := os.LookupEnv("CORS_ALLOWED_ORIGIN")
		var allowedOrigins []string

		if ok && allowedOrigin != "" {
			allowedOrigins = strings.Split(allowedOrigin, ",")
		} else {
			// 環境変数が見つからない場合はデフォルトのオリジンを指定
			allowedOrigins = []string{
				"http://localhost:3000",
				"http://localhost:5173",
				"https://todo-front-sable.vercel.app",
			}
		}

		origin := c.Request.Header.Get("Origin")
	
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Cookie や セッションを送信する場合は必須
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
