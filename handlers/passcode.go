package handlers

import (
	"net/http"
	"time"
	"strconv"

	"github.com/labstack/echo/v4"

	"api/db"
	"api/models"
	"api/responses"
	"api/messages"
	"api/middlewares"
)

// パスコード認証
func AdminPasscodes(c echo.Context) error {
	// リクエスト処理
	type Req struct {
		Passcode string `json:"passcode"`
	}
	req := new(Req)
	if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
    }

	id  := c.Param("id")
	uid := c.Param("uid")
	adminID, _ := strconv.ParseUint(uid, 10, 64)

	var passcode models.Passcode

	// idがなければ有効化失敗
	if err := db.DB.First(&passcode, id).Error; err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": messages.Status[7000]})
	}

	if passcode.AdminID != uint(adminID) {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[7001]})
	}

	if !middlewares.ComparePasscode(passcode.Code, req.Passcode) {
	    return c.JSON(http.StatusUnauthorized, echo.Map{"message": messages.Status[7002]})
	}

	// パスコードの有効期限チェック
	if passcode.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": messages.Status[7003]})
	}

	// ユーザー情報をactive
	if err := db.DB.Model(&models.Admin{}).
		Where("id = ?", uint(adminID)).
		Updates(map[string]interface{}{"status": "active"}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2006]})
	}
	// ユーザー情報を取得
	var admin models.Admin
	if err := db.DB.First(&admin, uint(adminID)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[4002]})
	}
	response := responses.Admin{
		ID:        admin.ID,
		Name:      admin.Name,
		Email:     admin.Email,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": messages.Status[1004],
		"user"   : response,
	})
}