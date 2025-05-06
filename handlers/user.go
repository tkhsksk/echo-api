package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"api/db"
	"api/models"
	"api/middlewares"
	"api/messages"
)

func GetUsers(c echo.Context) error {
	// デフォルト 100 件 上限1000（DoS対策）
	limit := middlewares.ParseLimitParam(c, 100, 1000)

	// db接続
	var users []models.User
	result := db.DB.Limit(limit).Find(&users)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2004]})
	}
	if result.RowsAffected == 0 {
	    return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4000]})
	}

	// 必要な情報だけをマッピング
	var response []models.UserResponse
	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1005],
		"users":   response,
	})
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")

	// db接続
	var user models.User
	result := db.DB.First(&user, id)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": messages.Status[4000]})
	}

	// 必要な情報だけをマッピング
	response := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1005],
		"user":    response,
	})
}

func GetUserProfile(c echo.Context) error {
	user := c.Get("user").(models.User)

	// 必要な情報だけをマッピング
	response := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1005],
		"user":    response,
	})
}
