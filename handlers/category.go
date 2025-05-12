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
func CreateCategory(c echo.Context) error {
	// ログイン中の管理者取得
	admin := c.Get("admin").(models.Admin)

	// 投稿内容を受け取る構造体
	type Req struct {
		Name     string  `json:"name" binding:"required"`
		Content  *string `json:"content"`
		Status   string  `json:"status" binding:"required"`
		ParentID *uint   `json:"parent_id"`
	}

	req := new(Req)
	// リクエストJSONを構造体にバインド
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	// 投稿データ作成
	category := models.Category{
		AdminID:  admin.ID,   	// ログイン管理者のID
		Name:     req.Name,   	// カテゴリー名
		Content:  req.Content,	// コンテンツ
		Status:   req.Status, 	// ステータス
		ParentID: req.ParentID, // 親id
	}

	// DBに保存
	if err := db.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2002]})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1001],
		"category": category,
	})
}

func GetCategories(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := middlewares.ParseLimitParam(c, 100, 1000)

	// db接続
	var categories []models.Category
	result := db.DB.Limit(limit).Find(&categories)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":    messages.Status[1005],
		"categories": categories,
	})
}

func GetCategoryByID(c echo.Context) error {
	id := c.Param("id")

	// db接続
	var category models.Category
	result := db.DB.First(&category, id)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1005],
		"category": result,
	})
}

func UpdateCategory(c echo.Context) error {
	// 投稿内容を受け取る構造体
	type Req struct {
		Name     string  `json:"name" binding:"required"`
		Content  *string  `json:"content"`
		Status   string  `json:"status" binding:"required"`
		ParentID *uint   `json:"parent_id"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil { return err } // 変換失敗時のエラー処理

	offset := idInt - 1
	if offset < 0 { offset = 0 }

	admin := c.Get("admin").(models.Admin)

	// db接続
	var category models.Category
	result := db.DB.First(&category, idInt)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	// データ更新
	category.AdminID  = admin.ID
	category.Name     = req.Name
	category.Content  = req.Content
	category.Status	  = req.Status
	category.ParentID = req.ParentID

	if err := db.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2005]})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1002],
		"category": category,
	})
}