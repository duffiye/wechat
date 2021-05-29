package subscribe

import (
	"encoding/json"
	"fmt"

	"github.com/duffiye/wechat/common"
	"github.com/duffiye/wechat/miniprogram/context"
)

const (
	subscribeMessage = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
)

type SubscribeRequest struct {
	ToUser           string      `json:"touser"`            //		是	接收者（用户）的 openid
	TemplateID       string      `json:"template_id"`       //	是	所需下发的订阅模板id
	Page             string      `json:"page"`              //		否	点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             interface{} `json:"data"`              //	是	模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
	MiniprogramState string      `json:"miniprogram_state"` //		否	跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string      `json:"lang"`              //	否	进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

type SubscribeResult struct {
	common.Error
}

type Subscribe struct {
	*context.Context
}

func NewSubscribe(ctx *context.Context) *Subscribe {
	return &Subscribe{ctx}
}

func (s *Subscribe) Send(request SubscribeRequest) (result SubscribeResult, err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return result, err
	}
	url := fmt.Sprintf(subscribeMessage, accessToken)
	fmt.Printf("url:%v\n", url)
	data, err := json.Marshal(&request)
	if err != nil {
		return SubscribeResult{}, err
	}
	fmt.Println(string(data))
	response, err := common.PostWithJSON(url, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Subscribe Message Send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return result, err
	}

	return result, err
}
