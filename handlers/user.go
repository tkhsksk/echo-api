package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"api/db"
	"api/models"
	"api/middlewares"
)

func GetUsers(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := middlewares.ParseLimitParam(c, 100, 1000)

	var users []models.User
	if err := db.DB.Limit(limit).Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "ユーザー取得に失敗しました"})
	}

	// 必要な情報だけをマッピング
	var response []models.UserResponse
	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:     u.ID,
			Email:  u.Email,
			Status: u.Status,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "ユーザーが見つかりません"})
	}

	// 必要な情報だけをマッピング
	response := models.UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Status: user.Status,
	}

	return c.JSON(http.StatusOK, response)
}
