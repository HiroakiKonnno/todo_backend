package db

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv" // godotenvライブラリを使う
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using default environment variables")
	}

	// フラグを定義
	debug := flag.Bool("debug", false, "デバッグモードを有効にします")
	flag.Parse()

	// 環境変数の取得
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || database == "" {
		panic("環境変数が不足しています。DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME を確認してください。")
	}

	if *debug {
		host = "localhost"
	}

	// DSNの作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		user, 
		password, 
		host, 
		port, 
		database,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	fmt.Printf("Connected to the MySQL database: %s\n", database)
}
