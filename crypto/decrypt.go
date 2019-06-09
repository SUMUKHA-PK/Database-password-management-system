package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// Decrypt decrypts
func Decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createMD5Hash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext, nil
}
