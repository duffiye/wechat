package wechatpay

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ZZMarquis/gm/util"
	"github.com/tjfoc/gmsm/sm4"
)

func TestJSAPI(t *testing.T) {

	out, err := sm4.Sm4Ecb([]byte("1234567890abcdef"), []byte("sddddd"), true)
	if err != nil {
		fmt.Println("我的错误信息")
		fmt.Println(err)
	}
	fmt.Println(out)
	fmt.Printf("%X\n", out)
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(out))
	fmt.Printf("%s", "======")

}

func TestHex(t *testing.T) {
	// data := `s`
	key := `583915df7a09dd0b9a644eda8e5ce5b2`
	str, err := hex.DecodeString(key)
	if err != nil {
		fmt.Printf("hex decode err %v: ", err)
	}

	fmt.Printf("%x\n", str)

}

func TestSm4ECBCipher_Encrypt(t *testing.T) {
	key := `583915df7a09dd0b9a644eda8e5ce5b2`
	hexDecodeKey, err := hex.DecodeString(key)
	if err != nil {
		fmt.Printf("hex decode err %v: ", err)
	}

	in := `s`
	c, err := sm4.NewCipher(hexDecodeKey)
	if err != nil {
		t.Error(err.Error())
		return
	}

	plainWithPadding := util.PKCS5Padding([]byte(in), sm4.BlockSize)
	result := make([]byte, len(plainWithPadding))
	c.Encrypt(result, plainWithPadding)

	fmt.Printf("%X\n", result)

	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%X", result))))
}
