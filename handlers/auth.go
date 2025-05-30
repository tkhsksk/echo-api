package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"

	"api/db"
	"api/models"
	"api/mailer"
	"api/messages"
	"api/utils"
)

// 管理者登録
func AdminRegister(c echo.Context) error {
	type Req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	if !utils.ValidateName(req.Name) {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[5002]})
	}

	if !utils.ValidateEmail(req.Email) {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[5000]})
	}

	if !utils.ValidatePassword(req.Password) {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[5001]})
	}

	// emailの重複チェック
	var adminCheck models.Admin
	if err := db.DB.Where("email = ?", req.Email).First(&adminCheck).Error; err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2007]})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	admin := models.Admin{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Status:   "suspended",
	}

	if err := db.DB.Create(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2002]})
	}

	// パスコードのハッシュ化
	passcode := utils.GenerateUnique6DigitCode()
	hashedPasscode, err := utils.HashPasscode(passcode)
	if err != nil {
	    utils.LogRequest(c, 004, "ハッシュ化失敗")
	}

	// パスコード作成
	passcode_model := models.Passcode{
		AdminID:   admin.ID,
		Code:	   hashedPasscode,
		ExpiresAt: time.Now().Add(1 * time.Hour), // 1時間有効
	}
	if err := db.DB.Create(&passcode_model).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2001]})
	}
	// 非同期メール送信
	go func() {
		err := mailer.SendPasscodeMail(req.Email, passcode, admin.ID, passcode_model.ID)
		if err != nil {
			utils.LogRequest(c, 002, "メール送信失敗")
		}
	}()

	return c.JSON(http.StatusCreated, echo.Map{"message": messages.Status[1003]})
}

// ユーザー登録
func UserRegister(c echo.Context) error {
	type Req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	if !utils.ValidateName(req.Name) {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[5002]})
	}

	if !utils.ValidateEmail(req.Email) {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[5000]})
	}

	if !utils.ValidatePassword(req.Password) {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[5001]})
	}

	// emailの重複チェック
	var userCheck models.User
	if err := db.DB.Where("email = ?", req.Email).First(&userCheck).Error; err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2007]})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Status:   "active",
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2002]})
	}

	utils.LogRequest(c, 003, "ユーザー登録成功")

	return c.JSON(http.StatusCreated, echo.Map{"message": messages.Status[1001]})
}

// ユーザーログイン
func UserLogin(c echo.Context) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": messages.Status[4000]})
	}

	if user.Status != "active" {
	    return c.JSON(http.StatusForbidden, echo.Map{"message": messages.Status[4001],})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": messages.Status[5001]})
	}

	// セッション作成
	session := models.UserSession{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(3 * time.Hour), // 3時間有効
	}

	if err := db.DB.Create(&session).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2001]})
	}

	// クッキーでセッションID返す
	c.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	})

	utils.LogRequest(c, 003, "ユーザーログイン成功")

	return c.JSON(http.StatusOK, echo.Map{
		"message":    messages.Status[1000],
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
		return c.JSON(http.StatusBadRequest, echo.Map{"message": messages.Status[2000]})
	}

	var admin models.Admin
	if err := db.DB.Where("email = ?", req.Email).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": messages.Status[4002]})
	}

	if admin.Status != "active" {
	    return c.JSON(http.StatusForbidden, echo.Map{"message": messages.Status[4003],})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": messages.Status[2003]})
	}

	// セッション作成
	session := models.AdminSession{
		ID:        uuid.New().String(),
		AdminID:   admin.ID,
		ExpiresAt: time.Now().Add(3 * time.Hour), // 3時間有効
	}

	if err := db.DB.Create(&session).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": messages.Status[2001]})
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
		"message":    messages.Status[1000],
		"session_id": session.ID,
	})
}
