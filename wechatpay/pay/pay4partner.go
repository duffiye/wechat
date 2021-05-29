package pay

import (
	"encoding/json"
	"fmt"

	cm "github.com/duffiye/wechat/common"

	"github.com/duffiye/wechat/wechatpay/common"
	"github.com/duffiye/wechat/wechatpay/context"

	"github.com/duffiye/wechat/wechatpay/pay/model"
)

const (
	// PartnerJSAPI 服务商模式-JSAPI下单
	PartnerJSAPI = "https://api.mch.weixin.qq.com/v3/pay/partner/transactions/jsapi"
)

// Pay4Partner 微信支付服务商模式
type Pay4Partner struct {
	*context.Context
}

// NewPay4Partner 创建微信支付服务商支付
func NewPay4Partner(ctx *context.Context) *Pay4Partner {
	return &Pay4Partner{ctx}
}

// JSAPIRequest jsapi 支付
type JSAPIRequest struct {
	common.Header

	// SpAppID 服务商公众号ID
	// string（32）	是	query 服务商户号，由微信支付生成并下发
	// 示例值：1230000109
	SpAppID string `json:"sp_appid"`

	// SpMchID 服务商户号
	// string（32）	是	query 服务商户号，由微信支付生成并下发
	// 示例值：1230000109
	SpMchID string `json:"sp_mchid"`

	// SubAppID 子商户公众号ID
	// string（32）	否	query 子商户申请的公众号或移动应用appid。
	// 示例值：wxd678efh567hg6999
	SubAppID string `json:"sub_appid"`

	// SubMchID 子商户号
	// string（32）	是	query 子商户的商户号，有微信支付生成并下发。
	// 示例值：1900000109
	SubMchID string `json:"sub_mchid"`

	// Description 商品描述
	// string（127）	是	query 商品描述
	// 示例值：Image形象店-深圳腾大-QQ公仔
	Description string `json:"description"`

	// OutTradeNo 商户订单号
	// string（32）	是	query 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一，详见【商户订单号】。
	// 特殊规则：最小字符长度为6
	// 示例值：1217752501201407033233368018
	OutTradeNo string `json:"out_trade_no"`

	// TimeExpire 交易结束时间
	// string（64）	否	query 订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。
	// 示例值：2018-06-08T10:34:56+08:00
	TimeExpire string `json:"time_expire"`

	// Attach 附加数据
	// string（128）	否	query 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	// 示例值：自定义数据
	Attach string `json:"attach"`

	// NotifyURL 通知地址
	// string（256）	是	query 通知URL必须为直接可访问的URL，不允许携带查询串。
	// 格式：URL
	// 示例值：https://www.weixin.qq.com/wxpay/pay.php
	NotifyURL string `json:"notify_url"`

	// GoodsTag 订单优惠标记
	// string（32）	否	query 订单优惠标记
	// 示例值：WXG
	GoodsTag string `json:"goods_tag"`

	// SettleInfo 结算信息
	// object	否	query 结算信息
	SettleInfo model.SettleInfo `json:"settle_info"`

	// Amount 订单金额
	// object	是	query 订单金额信息
	Amount model.Amount `json:"amount"`

	// Payer 支付者
	// object	是	query 支付者信息
	Payer model.Payer `json:"payer"`

	// Detail 优惠功能
	// object	否	query 优惠功能
	Detail model.Detail `json:"detail"`

	// SceneInfo 场景信息
	// object	否	query 支付场景描述
	SceneInfo model.SceneInfo `json:"scene_info"`
}

// JSAPIResult jsapi 返回
type JSAPIResult struct {

	// PrepayID 预支付交易会话标识
	// string（64）	是	预支付交易会话标识。
	// 示例值：wx201410272009395522657a690389285100
	PrepayID string `json:"prepay_id"`
}

// JSAPI jsapi 支付
func (p *Pay4Partner) JSAPI(request JSAPIRequest) (result JSAPIResult, err error) {
	var method string = "POST"
	var url string = "/v3/pay/partner/transactions/jsapi"
	requestBody, _ := json.Marshal(request)
	authorization, err := p.GetAuthorization(method, url, request.Timestamp, request.NonceStr, requestBody)
	if err != nil {
		return result, fmt.Errorf("获取Authorization信息失败,%v", err)
	}
	cm.PostWithAuthorizationSerial(PartnerJSAPI, requestBody, authorization, p.SerialNo)
	response, err := cm.PostWithAuthorizationSerial(PartnerJSAPI, requestBody, authorization, p.SerialNo)
	if err != nil {
		return result, err
	}
	fmt.Printf("服务商模式，JSAPI支付返回信息%s", response)

	fmt.Println(authorization)
	return
}
