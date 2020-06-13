package model

// SettleInfo  结算信息
type SettleInfo struct {

	// ProfitSharing 是否指定分账
	// 是否指定分账，枚举值 true：是 false：否
	// 示例值：true
	ProfitSharing bool `json:"profit_sharing"`

	// SubsidyAmount 补差金额
	// SettleInfo.profit_sharing为true时，该金额才生效。
	// 示例值：10
	SubsidyAmount int64 `json:"subsidy_amount"`
}
