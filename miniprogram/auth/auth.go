package auth

import (
	"encoding/json"
	"fmt"

	"github.com/duffiye/wechat/common"
	"github.com/duffiye/wechat/miniprogram/context"
)

const (
	jsCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

// Auth 小程序认证
type Auth struct {
	*context.Context
}

//NewAuth new auth
func NewAuth(ctx *context.Context) *Auth {
	return &Auth{ctx}
}

// Jscode2sessionResult 登录凭证校验返回信息
type Jscode2sessionResult struct {
	common.Error
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

// Jscode2session 登录
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (auth *Auth) Jscode2session(jsCode string) (result Jscode2sessionResult, err error) {
	urlStr := fmt.Sprintf(jsCode2SessionURL, auth.AppID, auth.AppSecret, jsCode)
	var response []byte
	response, err = common.Get(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return result, nil
}
