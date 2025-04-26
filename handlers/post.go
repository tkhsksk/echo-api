package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"api/db"
	"api/models"
	"api/middlewares"
	"api/messages"
)

// 投稿作成
func CreatePosts(c echo.Context) error {
	// ログイン中のユーザー取得
	user := c.Get("user").(models.User)

	// 投稿内容を受け取る構造体
	type Req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	// 投稿データ作成
	post := models.Post{
		UserID:  user.ID,     // ログインユーザーのID
		Title:   req.Title,   // タイトル
		Content: req.Content, // コンテンツ
	}

	// DBに保存
	if err := db.DB.Create(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2002]})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": messages.Status[1001], "post": post})
}

// 自身が登録した投稿一覧取得
func GetPosts(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := middlewares.ParseLimitParam(c, 100, 1000)
	user := c.Get("user").(models.User)

	// db接続
	var posts []models.Post
	// ユーザーidが一致しているものだけ取得
	result := db.DB.Limit(limit).Where("user_id = ?", user.ID).Find(&posts)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4004]})
	}

	// 必要な情報だけをマッピング
	var response []models.PostResponse
	for i, u := range posts {
		response = append(response, models.PostResponse{
			ID:        uint(i+1),
			Title:     u.Title,
			Content:   u.Content,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func GetPostByID(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil { return err } // 変換失敗時のエラー処理

	offset := idInt - 1
	if offset < 0 { offset = 0 }

	user := c.Get("user").(models.User)

	// db接続
	var posts models.Post
	result := db.DB.Where("user_id = ?", user.ID).Offset(offset).Limit(1).Find(&posts)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4004]})
	}

	// 必要な情報だけをマッピング
	response := models.PostResponse{
		ID:        uint(idInt),
		Title:     posts.Title,
		Content:   posts.Content,
		CreatedAt: posts.CreatedAt,
		UpdatedAt: posts.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response)
}