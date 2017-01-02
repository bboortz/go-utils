package utils

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	Seed()
}

func Seed() {
	rand.Seed(GenSeed())
	rand.Seed(int64(randomInt64(0, math.MaxInt64)))
}

func GenSeed() int64 {
	return time.Now().UTC().UnixNano() + int64(randomInt(0, 9999999))
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomInt64(min int64, max int64) int64 {
	return min + rand.Int63n(max-min)
}
