package main

import (
	"fmt"
	"todo_backend/internal/app/middleware"
	"todo_backend/internal/db"
	"todo_backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.New()
	r.Use(middleware.CORSMiddleware())    
	handler.RegisterTaskRoutes(r, db.DB)
	handler.RegisterAuthentificationRoutes(r, db.DB)

	port := "8080"
	fmt.Printf("サーバーがポート %s で起動中...\n", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("サーバーの起動に失敗しました: %v", err)
	}
}
