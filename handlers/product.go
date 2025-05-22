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
func CreateProduct(c echo.Context) error {
	// ログイン中の管理者取得
	admin := c.Get("admin").(models.Admin)

	// 投稿内容を受け取る構造体
	type Req struct {
		Name    string  `json:"name" binding:"required"`
		Price   float64 `json:"price" binding:"required"`
		Content *string `json:"content"`
		Status  string  `json:"status" binding:"required"`
		CategoryID uint `json:"category_id"`
	}

	req := new(Req)
	// リクエストJSONを構造体にバインド
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	// カテゴリーの存在チェック
	var category_str models.Category
	check := db.DB.First(&category_str, req.CategoryID)
	if check.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4006]})
	}

	// 投稿データ作成
	product := models.Product{
		AdminID: 	admin.ID,   	// ログイン管理者のID
		Name:    	req.Name,   	// 商品名
		Price:      req.Price,  	// 商品価格
		Content:    req.Content,	// コンテンツ
		Status:     req.Status, 	// ステータス
		CategoryID: req.CategoryID, // ステータス
	}

	// DBに保存
	if err := db.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2002]})
	}
	// Admin,Category取得用に再読み込み
	if err := db.DB.Preload("Admin").Preload("Category").First(&product, product.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}

	// 整形した構造体に詰める
	response := responses.NewProductForAdmin(product)

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1001],
		"product": response,
	})
}

func GetProductsForUser(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := utils.ParseLimitParam(c, 100, 1000)

	// db接続
	var products []models.Product
	result := db.DB.Limit(limit).Where("status = ?", "active").Find(&products)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	var response []responses.ProductForUser
    for _, p := range products {
        response = append(response, responses.ProductForUser{
            ID:    		p.ID,
            Name:  		p.Name,
            Price: 		p.Price,
            Content: 	*p.Content,
            CategoryID: p.CategoryID,
        })
    }

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1005],
		"products": response,
	})
}

func GetProductsForAdmin(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := utils.ParseLimitParam(c, 100, 1000)

	// db接続
	var products []models.Product
	result := db.DB.
		Preload("Category").
		Preload("Admin").
		Limit(limit).
		Find(&products)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	var response []responses.ProductForAdmin
	for _, p := range products {
		response = append(response, responses.NewProductForAdmin(p))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1005],
		"products": response,
	})
}

func GetProductsForAdminByCategory(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := utils.ParseLimitParam(c, 100, 1000)

	// db接続
	var products []models.Product
	result := db.DB.
		Preload("Category").
		Preload("Admin").
		Limit(limit).
		Find(&products)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	var response []responses.ProductForAdmin
	for _, p := range products {
		response = append(response, responses.NewProductForAdmin(p))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  messages.Status[1005],
		"products": response,
	})
}

func GetProductForUserByID(c echo.Context) error {
	id := c.Param("id")

	// db接続
	var product models.Product
	result := db.DB.
		Where("id = ? AND status = ?", id, "active").
		First(&product)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	// 整形した構造体に詰める
	response := responses.ProductForUser{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		CategoryID: product.Category.ID,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1005],
		"product": response,
	})
}

func GetProductForAdminByID(c echo.Context) error {
	id := c.Param("id")

	// db接続
	var product models.Product
	result := db.DB.
		Preload("Category").
		Preload("Admin").
		First(&product, id)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	// 整形した構造体に詰める
	response := responses.NewProductForAdmin(product)

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1005],
		"product": response,
	})
}

func UpdateProduct(c echo.Context) error {
	// 投稿内容を受け取る構造体
	type Req struct {
		Name    string  `json:"name" binding:"required"`
		Price   float64 `json:"price" binding:"required"`
		Content *string `json:"content"`
		Status  string  `json:"status" binding:"required"`
		CategoryID uint `json:"category_id"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	// カテゴリーの存在チェック
	var category_str models.Category
	check := db.DB.First(&category_str, req.CategoryID)
	if check.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4006]})
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil { return err } // 変換失敗時のエラー処理

	offset := idInt - 1
	if offset < 0 { offset = 0 }

	admin := c.Get("admin").(models.Admin)

	// db接続
	var product models.Product
	result := db.DB.First(&product, idInt)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4005]})
	}

	// データ更新
	product.AdminID    = admin.ID
	product.Name       = req.Name
	product.Price 	   = req.Price
	product.Content    = req.Content
	product.Status     = req.Status
	product.CategoryID = req.CategoryID

	if err := db.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2005]})
	}

	// Admin,Category取得用に再読み込み
	if err := db.DB.Preload("Admin").Preload("Category").First(&product, product.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}

	// 整形した構造体に詰める
	response := responses.NewProductForAdmin(product)

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1002],
		"product": response,
	})
}
