package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().Unix()))
}

// RandomInt randdom a number between min and max
func RandomInt(min, max int32) int32 {
	return rand.Int31n(max-min) + min
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

func RandomID() int32 {
	return rand.Int31n(899999999) + 100000000
}
