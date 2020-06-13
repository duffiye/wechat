package model

// StoreInfo 商户门店信息
// object	否	商户门店信息
type StoreInfo struct {

	// ID 门店ID
	// string（32）	否	商户侧门店编号
	// 示例值：0001
	ID string `json:"id"`

	// Name 门店名称
	// string（256）	是	商户侧门店名称
	// 示例值：腾讯大厦分店
	Name string `json:"name"`

	// area_code 地区编码
	// string（32）	是	地区编码，详细请见省市区编号对照表。
	// 示例值：440305
	AreaCode string `json:"area_code"`

	// Address 详细地址
	//	string（512）	是	详细的商户门店地址
	// 示例值：广东省深圳市南山区科技中一道10000号
	Address string `json:"address"`
}
