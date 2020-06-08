package context

import (
	"github.com/duffiye/wechat/credential"
	"github.com/duffiye/wechat/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
