package middlewares

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"api/db"
	"api/models"
)

// ログイン認証ミドルウェア（セッションIDをヘッダーから取得）
func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ヘッダーからセッションIDを取得
		sessionID := c.Request().Header.Get("Session-ID")
		if sessionID == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ログインしていません"})
		}

		// セッションIDが無効なら認証失敗
		var session models.Session
		if err := db.DB.Where("id = ?", sessionID).First(&session).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ログインセッションが無効です"})
		}

		// セッションの有効期限チェック
		if session.ExpiresAt.Before(time.Now()) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "セッションが期限切れです"})
		}

		// セッションが有効なら、ユーザー情報をコンテキストにセット
		var user models.User
		if err := db.DB.First(&user, session.UserID).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ユーザー情報が取得できません"})
		}

		// ユーザー情報を次の処理で使えるように設定
		c.Set("user", user)

		// 次の処理に進む
		return next(c)
	}
}

