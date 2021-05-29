package example

import (
	"fmt"
	"testing"

	"github.com/duffiye/wechat"
	"github.com/duffiye/wechat/wechatpay/config"
	"github.com/duffiye/wechat/wechatpay/marketing"
	"github.com/duffiye/wechat/wechatpay/pay"
	"github.com/duffiye/wechat/wechatpay/pay/model"
)

func TestFavorCallbacks(t *testing.T) {
	wc := wechat.NewWechat()
	cfg := &config.Config{}
	cfg.MchID = "mchid"
	cfg.SerialNo = "IJSINSF;XONA[XMDHF893NSX,SS[[SNDNFP[NSKNKSNI"
	wechatpay := wc.GetWechatPay(cfg)
	favorCallbacksRequest := marketing.FavorCallbacksRequest{}
	favorCallbacksRequest.Timestamp = "T22833843843"
	favorCallbacksRequest.NonceStr = "ssssssssssssss"
	favorCallbacksRequest.MchID = "229929292"
	result, err := wechatpay.GetMarketing().FavorCallbacks(favorCallbacksRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestJSAPI(t *testing.T) {
	wc := wechat.NewWechat()
	cfg := &config.Config{}
	cfg.MchID = "1308675601"
	cfg.SerialNo = "20AEACC22AD73745F1CFD6FB04E103058115EAC9"
	cfg.SpAppID = ""

	wechatpay := wc.GetWechatPay(cfg)
	request := pay.JSAPIRequest{}
	request.Amount = model.Amount{}

	pay := wechatpay.GetPay()
	result, err := pay.JSAPI(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}
