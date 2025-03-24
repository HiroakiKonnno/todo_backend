package db

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv" // godotenvライブラリを使う
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	env := os.Getenv("GO_ENV")

	if env != "prd" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("No .env file found, using default environment variables")
		}
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
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, password, database, port,
	)


	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	fmt.Printf("Connected to the MySQL database: %s\n", database)
}
