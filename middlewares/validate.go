package middlewares

import (
	"regexp"
)

// メールアドレスの正規表現
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// メールアドレスを検証する関数
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// 8文字以上20文字以内
// 英大文字（A-Z）を1文字以上
// 英小文字（a-z）を1文字以上
// 数字（0-9）を1文字以上
func ValidatePassword(pw string) bool {
	if len(pw) < 8 || len(pw) > 20 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasDigit := false

	for _, c := range pw {
		switch {
		case 'a' <= c && c <= 'z':
			hasLower = true
		case 'A' <= c && c <= 'Z':
			hasUpper = true
		case '0' <= c && c <= '9':
			hasDigit = true
		}
	}

	return hasLower && hasUpper && hasDigit
}