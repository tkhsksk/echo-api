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
    // ログ保存
    e.Use(middlewares.APILogMiddleware)

    // 画像登録
    e.File("/myfont.ttf", "static/myfont.ttf")
    e.File("/favicon.ico", "static/favicon.ico")
    e.File("/logo.svg", "static/logo.svg")
    e.File("/logo-dark.svg", "static/logo-dark.svg")

    // データベース削除
    e.POST("/delete", db.DbDelete)

	// 認証関連
	e.POST("/auth/admin/register", handlers.AdminRegister)
	e.POST("/auth/admin/login", handlers.AdminLogin)
	e.POST("/auth/user/register", handlers.UserRegister)
	e.POST("/auth/user/login", handlers.UserLogin)

	e.POST("/passcode/:id/:uid", handlers.AdminPasscodes)

	// 認証が必要なAPIにミドルウェアを適用
	// 通常ユーザー
	user := e.Group("/authed/user", middlewares.IsAuthenticatedUser)
	user.POST("/posts", handlers.CreatePost)
	user.PUT("/posts/:id", handlers.UpdatePost)
	user.GET("/posts", handlers.GetPosts)        // 一覧取得
	user.GET("/posts/:id", handlers.GetPostByID) // 個別取得

	// 商品
	user.GET("/products", handlers.GetProductsForUser) // 一覧取得

	user.GET("/profiles", handlers.GetUserProfile) // プロフィール取得

	// 管理者
	admin := e.Group("/authed/admin", middlewares.IsAuthenticatedAdmin)
	admin.GET("/users", handlers.GetUsers)        				  // 一覧取得
	admin.GET("/users/:id", handlers.GetUserByID) 				  // 個別取得

	admin.GET("/users/sessions", handlers.GetUserSessions) // 一覧取得
	admin.GET("/users/sessions/:id", handlers.GetSessionsByUserID) // 個別一覧取得

	// カテゴリー
	admin.POST("/categories", handlers.CreateCategory) // 作成
	admin.GET("/categories", handlers.GetCategories) // 一覧取得
	admin.GET("/categories/tree", handlers.GetCategoriesTree) // 一覧取得
	admin.GET("/categories/:id", handlers.GetCategoryByID) // 個別取得
	admin.PUT("/categories/:id", handlers.UpdateCategory) // 更新

	admin.POST("/products", handlers.CreateProduct) // 商品作成
	admin.GET("/products", handlers.GetProductsForAdmin) // 商品一覧取得
	admin.GET("/products/:id", handlers.GetProductByID) // 商品個別取得
	admin.PUT("/products/:id", handlers.UpdateProduct) // 商品更新

	admin.GET("/profiles", handlers.GetAdminProfile) // プロフィール取得

	e.Start(":4207")
}
