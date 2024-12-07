package main

import (
	"net/http"

	"backend/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())      // CORSの設定を適用
	r.Use(gin.Recovery())   
	r.GET("api/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.Run()
}
