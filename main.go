package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"api/handlers"
	"api/middlewares"

	"api/db"
)

func main() {
	db.Init()
	e := echo.New()

	// ルーティング登録
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// ログイン登録用
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	e.GET("/users", handlers.GetUsers)       // 一覧取得
	e.GET("/users/:id", handlers.GetUserByID) // 個別取得

	// 認証が必要なAPIにミドルウェアを適用
	r := e.Group("/posts", middlewares.IsAuthenticated)
	r.POST("/", handlers.CreatePost)

	e.Start(":4207")
}
