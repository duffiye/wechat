package context

import (
	"github.com/duffiye/wechat/credential"
	"github.com/duffiye/wechat/wechatpay/config"
)

type Context struct {
	*config.Config
	credential.AuthorizationHandle
}
