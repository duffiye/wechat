package analysis

import (
	"encoding/json"
	"fmt"

	"github.com/duffiye/wechat/common"
	"github.com/duffiye/wechat/miniprogram/context"
)

const (
	getDailyRetainURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=%s"
)

type Analysis struct {
	*context.Context
}

type DailyRetainRequest struct {
	BeginDate string `json:"begin_date"` // 开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   // 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
}

type DailyRetainResult struct {
	common.Error
	RefDate    string `json:"ref_date"`     // 日期
	VisitUVNew UV     `json:"visit_uv_new"` // 新增用户留存
	VisitUV    UV     `json:"visit_uv"`     // 活跃用户留存
}

type UV struct {
	Key   uint `json:"key"`
	Value uint `json:"value"`
}

func NewAnalysis(ctx *context.Context) *Analysis {
	return &Analysis{ctx}
}

func (analysis *Analysis) GetDailyRetain(request DailyRetainRequest) (result DailyRetainResult, err error) {
	var accessToken string
	accessToken, err = analysis.GetAccessToken()
	if err != nil {
		return result, err
	}
	url := fmt.Sprintf(getDailyRetainURL, accessToken)
	body, err := json.Marshal(request)
	if err != nil {
		return result, err
	}
	response, err := common.Post(url, "application/json", body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetDailyRetain error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return result, err
	}

	return result, err
}
