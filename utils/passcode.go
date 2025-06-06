package utils

import (
    "math/rand"
    "time"
    
    "golang.org/x/crypto/bcrypt"
)

// パスコードの生成
func GenerateUnique6DigitCode() string {
    rand.Seed(time.Now().UnixNano())
    code := ""
    counts := make(map[rune]int)

    for len(code) < 6 {
        d := rune('0' + rand.Intn(10))
        // この数字がすでに2回出ていたらスキップ
        if counts[d] >= 2 {
            continue
        }
        code += string(d)
        counts[d]++
    }
    return code
}

// パスコードをハッシュ化する関数
func HashPasscode(passcode string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(passcode), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashed), nil
}
