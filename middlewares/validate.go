package middlewares

import (
	"regexp"
)

// メールアドレスの正規表現
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// 8文字以上20文字以内
// 英大文字（A-Z）を1文字以上
// 英小文字（a-z）を1文字以上
// 数字（0-9）を1文字以上
var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,20}$`)

// メールアドレスを検証する関数
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// パスワードを検証する関数
func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)
}