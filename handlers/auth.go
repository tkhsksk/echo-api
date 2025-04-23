package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"

	"api/db"
	"api/models"
)

// 管理者登録
func AdminRegister(c echo.Context) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なリクエストです"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.Admin{
		Email:    req.Email,
		Password: string(hashed),
		Status:   "active",
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "登録に失敗しました"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "管理者登録完了"})
}

// ユーザー登録
func UserRegister(c echo.Context) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なリクエストです"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Email:    req.Email,
		Password: string(hashed),
		Status:   "active",
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "登録に失敗しました"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "ユーザー登録完了"})
}

// ユーザーログイン
func UserLogin(c echo.Context) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なリクエストです"})
	}

	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ユーザーが見つかりません"})
	}

	if user.Status != "active" {
	    return c.JSON(http.StatusForbidden, echo.Map{"error": "ユーザーがアクティブではありません",})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "パスワードが間違っています"})
	}

	// セッション作成
	session := models.UserSession{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(3 * time.Hour), // 3時間有効
	}

	if err := db.DB.Create(&session).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "セッション作成失敗"})
	}

	// クッキーでセッションID返す
	c.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"message": "ログイン成功",
		"session_id": session.ID,
	})
}

// 管理者ログイン
func AdminLogin(c echo.Context) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なリクエストです"})
	}

	var admin models.Admin
	if err := db.DB.Where("email = ?", req.Email).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "管理者が見つかりません"})
	}

	if admin.Status != "active" {
	    return c.JSON(http.StatusForbidden, echo.Map{"error": "管理者がアクティブではありません",})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "パスワードが間違っています"})
	}

	// セッション作成
	session := models.AdminSession{
		ID:        uuid.New().String(),
		AdminID:   admin.ID,
		ExpiresAt: time.Now().Add(3 * time.Hour), // 3時間有効
	}

	if err := db.DB.Create(&session).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "セッション作成失敗"})
	}

	// クッキーでセッションID返す
	c.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"message": "ログイン成功",
		"session_id": session.ID,
	})
}
