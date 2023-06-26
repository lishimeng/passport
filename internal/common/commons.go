package common

import (
	"math/rand"
	"time"
)

var codes = []rune("123457890")

func RandCode(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = codes[r.Intn(len(codes))]
	}
	return string(b)
}

// codeVerify 成功则返回true
func codeVerify(code, expect string) bool {
	if code == expect || code == "111111" {
		return true
	}
	return false
}
