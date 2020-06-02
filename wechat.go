package wechat

import (
	"github.com/duffiye/wechat/cache"
	"github.com/duffiye/wechat/miniprogram"
	miniConfig "github.com/duffiye/wechat/miniprogram/config"
)

// Wechat wc
type Wechat struct {
	cache cache.Cache
}

// NewWechat Create Wechat
func NewWechat() *Wechat {
	return &Wechat{}
}

//SetCache 设置cache
func (wc *Wechat) SetCache(cahce cache.Cache) {
	wc.cache = cahce
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}
