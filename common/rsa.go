package common

import (
	"crypto"
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
		return privateKey, fmt.Errorf("读取证书[%v]错误", err)
	}
	block, _ := pem.Decode(privateKeyBytes)
	privateInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return privateKey, fmt.Errorf("解析私钥信息错误%v", err)
	}
	return privateInterface.(*rsa.PrivateKey), nil
}

// SHA256withRSA SHA256withRSA
func SHA256withRSA(data string, pem string) (string, error) {
	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	privateKey, err := ParsePKCS1PrivateKey(pem)
	if err != nil {
		return "", fmt.Errorf("解析私钥正常错误[%v]，证书路径[%s]", err, pem)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", fmt.Errorf("生成签名信息错误[%v]", err)
	}

	fmt.Printf("Signature: %x\n", signature)
	signRet := fmt.Sprintf("%x", signature)
	fmt.Printf("sigRet: %s\n", signRet)

	return signRet, nil
}
