package common

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmonpqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RadomCrypto() string {
	randomIndex := rand.Intn(len(Cryptocurrencies))
	return Cryptocurrencies[randomIndex]
}

func RadomFiat() string {
	randomIndex := rand.Intn(len(FiatCurrencies))
	return FiatCurrencies[randomIndex]
}
