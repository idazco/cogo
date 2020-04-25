package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"
)

func Sha1OfString(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha1Random() string {
	var letters = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 20)
	for i := range b { b[i] = letters[rand.Intn(len(letters))] }
	return string(b)
}
