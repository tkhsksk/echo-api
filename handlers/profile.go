package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"api/db"
	"api/models"
	"api/responses"
	"api/messages"
	"api/utils"
)

func GetUserProfile(c echo.Context) error {
	user := c.Get("user").(models.User)

	// 必要な情報だけをマッピング
	response := responses.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.LogRequest(c, 003, "プロフィール取得成功")

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1005],
		"user":    response,
	})
}

func UpdateUserProfile(c echo.Context) error {
	// 投稿内容を受け取る構造体
	type Req struct {
		Name string `json:"name" binding:"required"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	user := c.Get("user").(models.User)
	user.Name = req.Name
	if err := db.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2005]})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1002],
		"user": user,
	})
}
