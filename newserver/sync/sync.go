package sync

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().Unix()))
}

func RandomID() int64 {
	return rand.Int63n(899999999) + 100000000
}
