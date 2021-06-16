package example

import (
	"fmt"
	"testing"

	"github.com/duffiye/wechat"
	"github.com/duffiye/wechat/miniprogram/auth"
	"github.com/duffiye/wechat/miniprogram/config"
	"github.com/duffiye/wechat/miniprogram/message/subscribe"

	"github.com/duffiye/wechat/cache"
)

func TestAuthCode2Session(t *testing.T) {
	//使用memcache保存access_token，也可选择redis或自定义cache
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     "",
		AppSecret: "",
		Cache:     memory,
	}
	result, err := wc.GetMiniProgram(cfg).GetAuth().Jscode2session("071H0Zsf2ErcTE0aDauf2gZIsf2H0Zs0")
	if err != nil {
		fmt.Printf("miniprogram code2session error %v\n", err)
	}
	fmt.Printf("Jscode2session Result %v\n", result)
}

func TestGetPaidUnionID(t *testing.T) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     "",
		AppSecret: "",
		Cache:     memory,
	}
	request := auth.GetPaidUnionIDRequest{}
	request.OpenID = "071H0Zsf2ErcTE0aDauf2gZIsf2H0Zs0"
	request.TransactionID = "sss"
	result, err := wc.GetMiniProgram(cfg).GetAuth().GetPaidUnionID(request)
	if err != nil {
		fmt.Printf("miniprogram GetPaidUnionID error %v\n", err)
	}
	fmt.Printf("GetPaidUnionID Result %v\n", result)
}

func TestSubscribeMessage(t *testing.T) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     "wx582c2a3ea6e645c8",
		AppSecret: "1a43bca4db711cfd5481ca41b0f94622",
		Cache:     memory,
	}
	request := subscribe.SubscribeRequest{}
	request.MiniprogramState = ""
	request.Page = "index"
	request.ToUser = "oyWPK5GDRA-KHZkUStK8OvoKqtcQ"
	request.Lang = "zh_CN"
	request.TemplateID = "OobKR3uFlqguH5b6xWkUJVu344FZwvw3tn1fv3xjs-s"
	request.Data = map[string]struct {
		Value string `json:"value"`
	}{
		"date3": struct {
			Value string "json:\"value\""
		}{Value: "2019年10月1日"},
		"thing4": struct {
			Value string "json:\"value\""
		}{
			Value: "sss",
		},
		"phrase5": struct {
			Value string "json:\"value\""
		}{
			Value: "待审核",
		},
	}
	ret, err := wc.GetMiniProgram(cfg).GetSubscribe().Send(request)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(ret)
	}
}
