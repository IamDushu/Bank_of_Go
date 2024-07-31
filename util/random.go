package util

import (
	"math/rand"
	"strings"
	"time"
)

var randSrc *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	randSrc = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + randSrc.Int63n(max-min+1)
}

// RandomString generates a randoom string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[randSrc.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random money value
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency name
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "INR"}
	n := len(currencies)
	return currencies[randSrc.Intn(n)]
}
