package sha

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func Decode(rawPass, salt string) string {
	key, _ := hex.DecodeString(salt)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(rawPass))
	base64Sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return base64Sha
}
