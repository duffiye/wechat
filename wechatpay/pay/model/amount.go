package model

// Amount 订单信息
type Amount struct {

	// Total 总金额
	// 订单总金额，单位为分。
	// 示例值：100
	Total int `json:"total"`

	// Currency 货币类型
	// CNY：人民币，境内商户号仅支持人民币。
	// 示例值：CNY
	Currency string `json:"currency"`
}
