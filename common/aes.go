package common

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// AESDecryptWithGCM AES-256-GCM
func AESDecryptWithGCM(data string, key string, nonceStr string, additionalData string) string {
	dataBytes, _ := base64.StdEncoding.DecodeString(data)
	block, err := aes.NewCipher([]byte(key))

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, _ := aesgcm.Open(nil, []byte(nonceStr), dataBytes, []byte(additionalData))

	return fmt.Sprintf("%s", plaintext)
}
