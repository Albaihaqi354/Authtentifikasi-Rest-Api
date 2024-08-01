package util

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(n int) string {
	var charsets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letters := make([]rune, n)
	for i := range letters {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charsets))))
		if err != nil {
			panic(err)
		}
		letters[i] = charsets[num.Int64()]
	}

	return string(letters)
}
