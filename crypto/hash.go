package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// CreateMD5Hash returns an MD5 hashed string
func CreateMD5Hash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// CreateSHA256Hash returns a SHA256 hash
func CreateSHA256Hash(key string) string {
	hashedData := sha256.New()
	//hash the data
	hashedData.Write([]byte(key))

	//convert the hashed byte data to hex string
	hashedString := hex.EncodeToString(hashedData.Sum(nil))

	return hashedString
}
