package random

import (
	"math/rand"
	"time"
)

var (
	random *rand.Rand
)

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Int 生成 max 以内的 int 随机数
func Int(max int) int {
	return random.Intn(max)
}

// IntBetween 生成 min 和 max 之间的随机数
func IntBetween(min, max int) int {
	rand := Int(max)
	for rand < min {
		rand = Int(max)
	}

	return rand
}

// String 生成指定长度的字符串
func String(lenth int) string {
	seed := "aAbBc0CdDeE1fFgG2hHiI3jJkK4lLmM5nNoO6pPqQr7RsSt8TuUvVw9WxXyYzZ"
	seedLen := len(seed)

	result := ""
	for i := 0; i < lenth; i++ {
		result += string(seed[Int(seedLen)-1])
	}

	return result
}
