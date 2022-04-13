package wallet

import (
	"math/rand"
	"time"
)

const (
	letters           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	certificateLength = 32
)

func CreateRandomPassword() string {
	return randStringBytes(certificateLength)
}

func randStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
