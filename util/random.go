package util

import (
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt() int {
	return r.Int()
}

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[r.Intn(len(alphabet))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(10)
}

func RandomMoney() int64 {
	return int64(r.Intn(1000))
}

func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", RandomString(10))
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
