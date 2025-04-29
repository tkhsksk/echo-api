package middlewares

import (
	"math/rand"
	"time"
)

func GenerateUnique4DigitCode() string {
	rand.Seed(time.Now().UnixNano())

	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})

	// 先頭0もOKなのでそのまま先頭から4桁
	return string(digits[:4])
}
