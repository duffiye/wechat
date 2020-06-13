package model

// Payer 支付者
type Payer struct {

	// SpOpenID 用户服务标识
	// 是	用户在服务商appid下的唯一标识。
	// 示例值：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o
	SpOpenID string `json:"sp_openid"`

	// SubOpenID 用户子标识
	// 用户在子商户appid下的唯一标识。
	// 示例值：oUpF8uMuAJO_M2pxb1Q9zNjWeS6o
	SubOpenID string `json:"sub_openid"`
}
