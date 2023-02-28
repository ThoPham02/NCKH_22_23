package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().Unix()))
}

// RandomInt randdom a number between min and max
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// RandomString random string having n characters
func RandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	len := len(letters)
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len)]
	}
	return string(b)
}
