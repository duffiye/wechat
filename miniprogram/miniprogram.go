package miniprogram

import (
	"github.com/duffiye/wechat/credential"
	"github.com/duffiye/wechat/miniprogram/auth"
	"github.com/duffiye/wechat/miniprogram/config"
	"github.com/duffiye/wechat/miniprogram/context"
)

// MiniProgram 小程序相关
type MiniProgram struct {
	ctx *context.Context
}

// NewMiniProgram new miniprogram
func NewMiniProgram(cfg *config.Config) *MiniProgram {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)

	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &MiniProgram{ctx}
}

// GetContext get Context
func (miniProgram *MiniProgram) GetContext() *context.Context {
	return miniProgram.ctx
}

//GetAuth 登录/用户信息相关接口
func (miniProgram *MiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.ctx)
}
