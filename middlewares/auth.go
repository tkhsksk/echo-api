package middlewares

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"api/db"
	"api/models"
)

// ユーザーログイン認証ミドルウェア（セッションIDをpostから取得）
func IsAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ヘッダーからセッションIDを取得
		sessionID := c.Request().Header.Get("Session-ID")
		if sessionID == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ログインしていません"})
		}

		// セッションIDが無効なら認証失敗
		var session models.UserSession
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

		if user.Status != "active" {
		    return c.JSON(http.StatusForbidden, echo.Map{"error": "ユーザーがアクティブではありません",})
		}

		// ユーザー情報を次の処理で使えるように設定
		c.Set("user", user)

		// 次の処理に進む
		return next(c)
	}
}

// 管理者ログイン認証ミドルウェア（セッションIDをpostから取得）
func IsAuthenticatedAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ヘッダーからセッションIDを取得
		sessionID := c.Request().Header.Get("Session-ID")
		if sessionID == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ログインしていません"})
		}

		// セッションIDが無効なら認証失敗
		var session models.AdminSession
		if err := db.DB.Where("id = ?", sessionID).First(&session).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ログインセッションが無効です"})
		}

		// セッションの有効期限チェック
		if session.ExpiresAt.Before(time.Now()) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "セッションが期限切れです"})
		}

		// セッションが有効なら、ユーザー情報をコンテキストにセット
		var admin models.Admin
		if err := db.DB.First(&admin, session.AdminID).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "管理者情報が取得できません"})
		}

		if admin.Status != "active" {
		    return c.JSON(http.StatusForbidden, echo.Map{"error": "管理者がアクティブではありません",})
		}

		// ユーザー情報を次の処理で使えるように設定
		c.Set("admin", admin)

		// 次の処理に進む
		return next(c)
	}
}

