package auth

import (
	"encoding/json"
	"fmt"

	"github.com/duffiye/wechat/common"
	"github.com/duffiye/wechat/miniprogram/context"
)

const (
	jsCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	getPaidUnionIDURL = "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s"
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

// GetPaidUnionIDResult 获取用户支付后UnionID返回信息
type GetPaidUnionIDResult struct {
	common.Error
	UnionID string
}

// GetPaidUnionIDRequest 获取用户支付后UnionID请求信息
type GetPaidUnionIDRequest struct {
	OpenID        string
	TransactionID string
	MchID         string
	OutTradeNo    string
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

// GetPaidUnionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (auth *Auth) GetPaidUnionID(request GetPaidUnionIDRequest) (result GetPaidUnionIDResult, err error) {
	var accessToken string
	accessToken, err = auth.GetAccessToken()
	if err != nil {
		return result, err
	}
	url := fmt.Sprintf(getPaidUnionIDURL, accessToken, request.OpenID)
	if common.IsNotBlank(request.TransactionID) {
		url += fmt.Sprintf("&transaction_id=%s", request.TransactionID)
	}
	if common.IsNotBlank(request.OutTradeNo) && common.IsNotBlank(request.MchID) {
		url += fmt.Sprintf("&out_trade_no=%s&mch_id=%s", request.OutTradeNo, request.MchID)
	}
	response, err := common.Get(url)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return result, err
	}

	return result, err
}
