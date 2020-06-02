package context

import (
	"github.com/duffiye/wechat/miniprogram/config"
	// "github.com/silenceper/wechat/v2/credential"
)

// Context struct
type Context struct {
	*config.Config
	// credential.AccessTokenHandle
}
