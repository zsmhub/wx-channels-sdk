package apis

import (
	"encoding/json"
)

// 获取提现记录列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/getwithdrawlist.html

type ReqFundsGetwithdrawlist struct {
	EndTime   int `json:"end_time"`
	PageNum   int `json:"page_num"`
	PageSize  int `json:"page_size"`
	StartTime int `json:"start_time"`
}

var _ bodyer = ReqFundsGetwithdrawlist{}

func (x ReqFundsGetwithdrawlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsGetwithdrawlist struct {
	CommonResp
	TotalNum    int      `json:"total_num"`
	WithdrawIds []string `json:"withdraw_ids"`
}

var _ bodyer = RespFundsGetwithdrawlist{}

func (x RespFundsGetwithdrawlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsGetwithdrawlist(req ReqFundsGetwithdrawlist) (RespFundsGetwithdrawlist, error) {
	var resp RespFundsGetwithdrawlist
	err := c.executeWXApiPost("/channels/ec/funds/getwithdrawlist", req, &resp, true)
	if err != nil {
		return RespFundsGetwithdrawlist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsGetwithdrawlist{}, bizErr
	}
	return resp, nil
}
