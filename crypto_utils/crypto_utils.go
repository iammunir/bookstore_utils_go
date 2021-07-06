package cryptoutils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(password string) string {

	hash := md5.New()
	hash.Write([]byte(password))
	defer hash.Reset()

	return hex.EncodeToString(hash.Sum(nil))
}
