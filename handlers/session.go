package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"api/db"
	"api/models"
	"api/middlewares"
)

func GetUserSessions(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := middlewares.ParseLimitParam(c, 100, 1000)

	var sessions []models.UserSession
	if err := db.DB.Limit(limit).Find(&sessions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "セッション取得に失敗しました"})
	}

	// 必要な情報だけをマッピング
	var response []models.UserSessionResponse
	for _, u := range sessions {
		response = append(response, models.UserSessionResponse{
			ID:        u.ID,
			UserID:    u.UserID,
			CreatedAt: u.CreatedAt,
			ExpiresAt: u.ExpiresAt,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func GetSessionsByUserID(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := middlewares.ParseLimitParam(c, 100, 1000)
	// urlからid取得
	id := c.Param("id")
	var sessions []models.UserSession
	if err := db.DB.Limit(limit).Where("user_id = ?", id).Find(&sessions).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "セッションが見つかりません"})
	}

	// 必要な情報だけをマッピング
	var response []models.UserSessionResponseByUserID
	for _, u := range sessions {
		response = append(response, models.UserSessionResponseByUserID{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			ExpiresAt: u.ExpiresAt,
		})
	}

	return c.JSON(http.StatusOK, response)
}
