package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// EncryptOAEP 加密
func EncryptOAEP(text string, pemPath string) string {
	rsaPublicKey, err := ParsePKIXPublicKey(pemPath)
	fmt.Printf("读取证书[%s]错误[%v]", pemPath, err)
	secretMessage := []byte(text)
	rng := rand.Reader
	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rng, rsaPublicKey, secretMessage, nil)
	if err != nil {
		return ""
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	return ciphertext
}

// DecryptOAEP 解密
func DecryptOAEP(ciphertext string, pemPath string) string {
	rsaPrivateKey, _ := ParsePKCS1PrivateKey(pemPath)
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, rsaPrivateKey, cipherdata, nil)
	if err != nil {
		return ""
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	return string(plaintext)
}

// ParsePKIXPublicKey 获取RSA公钥
func ParsePKIXPublicKey(pemPath string) (publicKey *rsa.PublicKey, err error) {
	publicKeyBytes, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return publicKey, err
	}
	block, _ := pem.Decode(publicKeyBytes)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return publicKey, err
	}
	return pubInterface.(*rsa.PublicKey), nil
}

// ParsePKCS1PrivateKey 解析私钥
func ParsePKCS1PrivateKey(pemPath string) (privateKey *rsa.PrivateKey, err error) {
	privateKeyBytes, err := ioutil.ReadFile(pemPath)
	if err != nil {
		fmt.Println(err)
		return privateKey, err
	}
	block, _ := pem.Decode(privateKeyBytes)
	privateInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return privateKey, err
	}
	return privateInterface, nil
}
