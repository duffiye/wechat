package wechatpay

import (
	"github.com/duffiye/wechat/credential"
	"github.com/duffiye/wechat/wechatpay/config"
	"github.com/duffiye/wechat/wechatpay/context"

	"github.com/duffiye/wechat/wechatpay/marketing"
)

type WechatPay struct {
	*context.Context
}

func NewWechatPay(cfg *config.Config) *WechatPay {
	authorizationHandle := credential.NewDefaultAuthorization(cfg.MchID, cfg.SerialNo)
	ctx := &context.Context{
		Config:              cfg,
		AuthorizationHandle: authorizationHandle,
	}
	return &WechatPay{ctx}
}

func (wp *WechatPay) GetMarketing() *marketing.Marketing {
	return marketing.NewMarketing(wp.Context)
}
