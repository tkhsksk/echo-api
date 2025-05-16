package middlewares

import (
	"golang.org/x/crypto/bcrypt"
)

// 入力とハッシュの照合
func ComparePasscode(hashedPasscode, input string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPasscode), []byte(input))
    return err == nil
}
