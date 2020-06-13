package model

// SceneInfo 场景信息
// object	否	query 支付场景描述
type SceneInfo struct {

	// PayerClientIP 用户终端IP	
	// string（45）	是	调用微信支付API的机器IP，支持IPv4和IPv4两种格式的IP地址。
	// 示例值：14.23.150.211
	PayerClientIP string `json:"payer_client_ip"`

	// DeviceID 商户端设备号	
	// string（32）	否	商户端设备号（门店号或收银设备ID）。
	// 示例值：013467007045764
	DeviceID string `json:"device_id"`

	// StoreInfo 商户门店信息
	// object	否	商户门店信息
	StoreInfo StoreInfo `json:"store_info"`
}