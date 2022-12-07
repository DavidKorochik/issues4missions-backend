package util

import (
	"math/rand"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrctuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(limit int) (randomString string) {
	randomIndex := rand.Intn(limit)

	for i := 0; i < limit; i++ {
		randomString += string(alphabet[randomIndex])
	}

	return
}

func RandomTitle() string {
	return RandomString(8)
}

func RandomDescription() string {
	return RandomString(20)
}
