package main

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
	"api/handlers"
	"api/middlewares"

	"api/db"
)

// テンプレートレンダラー定義
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	db.Init()
	e := echo.New()

	// テンプレートのセットアップ
    t := &Template{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
    e.Renderer = t

	// ルーティング登録
	e.GET("/", func(c echo.Context) error {
        data := map[string]interface{}{
            "Title": "api.ksk318.me",
        }
        return c.Render(http.StatusOK, "index.html", data)
    })

    // 画像登録
    e.File("/favicon.ico", "static/favicon.ico")
    e.File("/logo.svg", "static/logo.svg")
    e.File("/logo-dark.svg", "static/logo-dark.svg")

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
