package credential

import (
	"fmt"
)

// AuthorizationHandle 获取微信支付的认证信息
type AuthorizationHandle interface {
	GetAuthorization(method, url, timestamp, nonceStr string, body []byte) (authorization string, err error)
}

// DefaultAuthorization 微信支付的认证信息
type DefaultAuthorization struct {
	SerialNo string
	MchID    string
}

// NewDefaultAuthorization 创建微信支付的认证信息
func NewDefaultAuthorization(mchID, serialNo string) AuthorizationHandle {
	return &DefaultAuthorization{
		MchID:    mchID,
		SerialNo: serialNo,
	}
}

// GetAuthorization 获取认证信息
func (da *DefaultAuthorization) GetAuthorization(method, url, timestamp, nonceStr string, body []byte) (authorization string, err error) {
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", method, url, timestamp, nonceStr, string(body))
	signature := GetSignature(message)
	authorization = fmt.Sprintf("%s mchid=\"%s\",serial_no=\"%s\",nonce_str=\"%s\",timestamp=\"%s\",signature=\"%s\"", "WECHATPAY2-SHA256-RSA2048", da.MchID, da.SerialNo, nonceStr, timestamp, signature)
	return authorization, nil
}

// GetSignature 获取签名信息
func GetSignature(body string) string {
	return body
}
