package common

import (
	"fmt"
	"testing"
)

func TestEncryptOAEP(t *testing.T) {
	var text string = "sss"
	var pemPath string = "sss"
	plainText := EncryptOAEP(text, pemPath)
	fmt.Printf("加密后的内容%v\n", plainText)
}
