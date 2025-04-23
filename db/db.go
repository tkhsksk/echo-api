package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

	"api/models"
)

var DB *gorm.DB

func Init() {
	// .env を読み込む
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 例: .env または os.Getenv で接続情報を取得（環境変数がベスト）
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// DSN（Data Source Name）を組み立てる
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	// DB接続
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}

	// モデルのマイグレーション（初回起動時などにテーブル自動作成）
	if err := DB.AutoMigrate(
		&models.Admin{},
		&models.AdminSession{},
		&models.User{},
		&models.UserSession{},
		&models.Post{},
	); err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}
}
