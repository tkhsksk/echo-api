package db

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"api/messages"
)

var ClearDB *gorm.DB

func DbDelete(c echo.Context) error {
	// リクエスト処理
	type Req struct {
		Password string `json:"password"`
	}
	req := new(Req)
	if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
    }

    // パスワード認証
    if os.Getenv("DB_DELETE_PASS") != req.Password {
	    return c.JSON(http.StatusUnauthorized, echo.Map{"message": messages.Status[2003],})
	}

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
	ClearDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}

	// 外部キー制約を無効化
	ClearDB.Exec("SET foreign_key_checks = 0;")

	// TRUNCATE コマンドを使って各テーブルをリセット
	tables := []string{
		"admins",
		"admin_sessions",
		"users",
		"user_sessions",
		"posts",
		"passcodes",
		"api_logs",
	}

	// 各テーブルに対して TRUNCATE を実行
	for _, table := range tables {
		result := ClearDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
		if result.Error != nil {
			log.Fatalf("テーブル %s の TRUNCATE に失敗しました: %v", table, result.Error)
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": fmt.Sprintf("%s の TRUNCATE に失敗しました", table)})
		}
	}

	// 外部キー制約を再度有効化
	ClearDB.Exec("SET foreign_key_checks = 1;")

	// 成功レスポンス
	return c.JSON(http.StatusOK, echo.Map{
		"message": "データベースのテーブルがリセットされました",
	})
}
