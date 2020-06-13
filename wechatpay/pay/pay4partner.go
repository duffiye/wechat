package pay

import (
	"github.com/duffiye/wechat/wechatpay/context"
)

const (
	// PartnerJSAPI 服务商模式-JSAPI下单
	PartnerJSAPI = "https://api.mch.weixin.qq.com/v3/pay/partner/transactions/jsapi"
)

// Pay4Partner 微信支付服务商模式
type Pay4Partner struct {
	*context.Context
}

// NewPay4Partner 创建微信支付服务商支付
func NewPay4Partner(ctx *context.Context) *Pay4Partner {
	return &Pay4Partner{ctx}
}
