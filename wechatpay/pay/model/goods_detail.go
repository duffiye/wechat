package model

// GoodsDetail 单品列表
// array	否	单品列表信息
// 条目个数限制：【1，undefined】
type GoodsDetail struct {

	// MerchantGoodsID 商户侧商品编码
	// string（32）	是	由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	// 示例值：商品编码
	MerchantGoodsID string `json:"merchant_goods_id"`

	// WechatpayGoodsID 微信侧商品编码
	// string（32）	否	微信支付定义的统一商品编号（没有可不传）
	// 示例值：1001
	WechatpayGoodsID string `json:"wechatpay_goods_id"`

	// GoodsName 商品名称
	// 	string（256）	否	商品的实际名称
	// 示例值：iPhoneX 256G
	GoodsName string `json:"goods_name"`

	// Quantity 商品数量
	// int	是	用户购买的数量
	// 示例值：1
	Quantity int `json:"quantity"`

	// UnitPrice 商品单价
	// int	是	商品单价，单位为分
	// 示例值：828800
	UnitPrice int `json:"unit_price"`
}
