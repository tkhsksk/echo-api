package middlewares

import (
	"math/rand"
	"time"
)

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
