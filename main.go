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
    e.File("/myfont.ttf", "static/myfont.ttf")
    e.File("/favicon.ico", "static/favicon.ico")
    e.File("/logo.svg", "static/logo.svg")
    e.File("/logo-dark.svg", "static/logo-dark.svg")

	// 認証関連
	e.POST("/auth/admin/register", handlers.AdminRegister)
	e.POST("/auth/admin/login", handlers.AdminLogin)
	e.POST("/auth/user/register", handlers.UserRegister)
	e.POST("/auth/user/login", handlers.UserLogin)

	// 認証が必要なAPIにミドルウェアを適用
	userAuthGroup := e.Group("/posts", middlewares.IsAuthenticatedUser)
	userAuthGroup.POST("", handlers.CreatePosts)
	userAuthGroup.GET("", handlers.GetPosts)        // 一覧取得
	userAuthGroup.GET("/:id", handlers.GetPostByID)        // 個別取得

	adminAuthGroup := e.Group("/users", middlewares.IsAuthenticatedAdmin)
	adminAuthGroup.GET("", handlers.GetUsers)        // 一覧取得
	adminAuthGroup.GET("/:id", handlers.GetUserByID) // 個別取得
	adminAuthGroup.GET("/sessions", handlers.GetUserSessions) // 一覧取得
	adminAuthGroup.GET("/sessions/:id", handlers.GetSessionsByUserID) // 個別一覧取得

	e.Start(":4207")
}
