package lib

import (
	"math/rand"
	"time"
)

const (
	size      = 10
	rndSource = "BCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func GenerateUrlAlias() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune(rndSource)
	urlAlias := make([]rune, size)

	for i := range urlAlias {
		urlAlias[i] = chars[rnd.Intn(len(chars))]
	}

	return string(urlAlias)
}
