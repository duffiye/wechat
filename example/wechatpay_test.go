package example

import (
	"fmt"
	"testing"

	"github.com/duffiye/wechat"
	"github.com/duffiye/wechat/wechatpay/config"
	"github.com/duffiye/wechat/wechatpay/marketing"
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
