package main

import (
	"fmt"
	"todo_backend/internal/db"
	"todo_backend/internal/handler"
	"todo_backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.New()
	r.Use(middleware.CORSMiddleware()) 

	public := r.Group("/")
	handler.RegisterAuthentificationRoutes(public, db.DB)


	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	handler.RegisterTaskRoutes(protected, db.DB)
	

	port := "8080"
	fmt.Printf("サーバーがポート %s で起動中...\n", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("サーバーの起動に失敗しました: %v", err)
	}
}
