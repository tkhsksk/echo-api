package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"api/db"
	"api/models"
)

// 投稿作成
func CreatePost(c echo.Context) error {
	// ログイン中のユーザー取得
	user := c.Get("user").(models.User)

	// 投稿内容を受け取る構造体
	type Req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// 投稿データ作成
	post := models.Post{
		UserID:  user.ID,    // ログインユーザーのID
		Title:   req.Title,   // タイトル
		Content: req.Content, // コンテンツ
	}

	// DBに保存
	if err := db.DB.Create(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "投稿の保存に失敗しました"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "投稿が作成されました", "post": post})
}
