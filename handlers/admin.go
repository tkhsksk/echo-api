package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	// "api/db"
	"api/models"
	"api/responses"
	// "api/middlewares"
	"api/messages"
)

func GetAdminProfile(c echo.Context) error {
	admin := c.Get("admin").(models.Admin)

	// 必要な情報だけをマッピング
	response := responses.Admin{
		ID:        admin.ID,
		Name:      admin.Name,
		Email:     admin.Email,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": messages.Status[1005],
		"admin":   response,
	})
}
