package marketing

import (
	"encoding/json"
	"fmt"

	"github.com/duffiye/wechat/wechatpay/common"

	"github.com/duffiye/wechat/wechatpay/context"
)

type Marketing struct {
	*context.Context
}

func NewMarketing(ctx *context.Context) *Marketing {

	return &Marketing{ctx}
}

type FavorCallbacksRequest struct {
	common.Header
	MchID     string `json:"mchid"`
	NotifyURL string `json:"notify_url"`
	Switch    bool   `json:"switch"`
}

type FavorCallbacksResult struct {
	UpdateTime string `json:"update_time"`
	NotifyURL  string `json:"notify_url"`
}

func (m *Marketing) FavorCallbacks(request FavorCallbacksRequest) (result FavorCallbacksResult, err error) {
	var method string = "POST"
	var url string = "/test"
	requestBody, _ := json.Marshal(request)
	authorization, err := m.GetAuthorization(method, url, request.Timestamp, request.NonceStr, requestBody)
	if err != nil {
		return result, fmt.Errorf("获取Authorization信息失败,%v", err)
	}
	fmt.Println(authorization)
	return
}
