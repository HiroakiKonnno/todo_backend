package main

import (
	"net/http"

	"todo_backend/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())      // CORSの設定を適用        
	r.GET("api/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.Run()
}
