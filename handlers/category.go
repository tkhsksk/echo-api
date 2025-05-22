package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"api/db"
	"api/models"
	"api/responses"
	"api/utils"
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
		// c.Logger().Errorf("Bind error: %v", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	// 親カテゴリーの存在チェック
	var category_str models.Category
	check := db.DB.First(&category_str, req.ParentID)
	if req.ParentID != nil && check.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4007]})
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

	// Admin取得用に読み込み
	if err := db.DB.Preload("Admin").First(&category, category.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}

	// api用に整形した構造体に詰める
	response := responses.Category{
		ID:       category.ID,
		Name:     category.Name,
		Status:   category.Status,
		ParentID: category.ParentID,
		Admin: responses.AdminSummary{
			ID:     category.Admin.ID,
			Name:   category.Admin.Name,
			Status: category.Admin.Status,
		},
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1001],
		"category": response,
	})
}

func GetCategories(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := utils.ParseLimitParam(c, 100, 1000)

	// db接続
	var categories []models.Category
	result := db.DB.Preload("Admin").Limit(limit).Find(&categories)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	var response []responses.Category
	for _, c := range categories {
		response = append(response, responses.Category{
			ID:       c.ID,
			Name:     c.Name,
			Status:   c.Status,
			ParentID: c.ParentID,
			Admin: responses.AdminSummary{
				ID:     c.Admin.ID,
				Name:   c.Admin.Name,
				Status: c.Admin.Status,
			},
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":    messages.Status[1005],
		"categories": response,
	})
}

func GetCategoriesTree(c echo.Context) error {
    var categories []models.Category
    if err := db.DB.Find(&categories).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
    }

    tree := buildCategoryTree(categories, nil)

    return c.JSON(http.StatusOK, echo.Map{
        "message":    messages.Status[1005],
        "categories": tree,
    })
}

// 再帰的にカテゴリーのツリーを構築
func buildCategoryTree(categories []models.Category, parentID *uint) []responses.CategoryTree {
    var tree []responses.CategoryTree
    for _, category := range categories {
        if (category.ParentID == nil && parentID == nil) || (category.ParentID != nil && parentID != nil && *category.ParentID == *parentID) {
            // 再帰的に子供を取得
            children := buildCategoryTree(categories, &category.ID)
            node := responses.CategoryTree{
                ID:       category.ID,
                Name:     category.Name,
                Status:   category.Status,
                Children: children,
            }
            tree = append(tree, node)
        }
    }
    return tree
}

func GetCategoryByID(c echo.Context) error {
	id := c.Param("id")

	// 指定されたカテゴリ取得
	var category models.Category
	result := db.DB.Preload("Admin").First(&category, id)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4006]})
	}

	// 親カテゴリを辿る
	var parentChain []map[string]interface{}
	currentParentID := category.ParentID

	for currentParentID != nil {
		var parent models.Category
		if err := db.DB.First(&parent, *currentParentID).Error; err != nil {
			break // 途中で見つからなければ終了
		}

		parentChain = append([]map[string]interface{}{ // prepend する
			{
				"id":   parent.ID,
				"name": parent.Name,
			},
		}, parentChain...)

		currentParentID = parent.ParentID
	}

	// レスポンス構造体に詰める
	response := responses.Category{
		ID:       category.ID,
		Name:     category.Name,
		Status:   category.Status,
		ParentID: category.ParentID,
		Admin: responses.AdminSummary{
			ID:     category.Admin.ID,
			Name:   category.Admin.Name,
			Status: category.Admin.Status,
		},
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":          messages.Status[1005],
		"category":         response,
		"parent_hierarchy": parentChain,
	})
}


func UpdateCategory(c echo.Context) error {
	// 投稿内容を受け取る構造体
	type Req struct {
		Name     string  `json:"name" binding:"required"`
		Content  *string `json:"content"`
		Status   string  `json:"status" binding:"required"`
		ParentID *uint   `json:"parent_id"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	// ParentIDが存在する時だけ親カテゴリーの存在チェック
	var category_str models.Category
	check := db.DB.First(&category_str, req.ParentID)
	if req.ParentID != nil && check.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4007]})
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

	// レスポンス構造体に詰める
	response := responses.Category{
		ID:       category.ID,
		Name:     category.Name,
		Status:   category.Status,
		ParentID: category.ParentID,
		Admin: responses.AdminSummary{
			ID:     admin.ID,
			Name:   admin.Name,
			Status: admin.Status,
		},
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1002],
		"category": response,
	})
}