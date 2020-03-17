package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5OfString(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func MD5OfBytes(text []byte) []byte {
	hasher := md5.New()
	hasher.Write(text)
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}
