package lib

import (
	"math/rand"
	"os"
	"time"
)

const (
	size      = 10
	rndSource = "RND_SOURCE" //"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func GenerateUrlAlias() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune(os.Getenv(rndSource))
	urlAlias := make([]rune, size)

	for i := range urlAlias {
		urlAlias[i] = chars[rnd.Intn(len(chars))]
	}

	return string(urlAlias)
}
