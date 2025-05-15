package apis

import (
	"encoding/json"
)

// 获取资金流水列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/getfundsflowlist.html

type ReqFundsGetfundsflowlist struct {
	EndTime       int    `json:"end_time"`
	NextKey       string `json:"next_key"`
	Page          int    `json:"page"`
	PageSize      int    `json:"page_size"`
	StartTime     int    `json:"start_time"`
	TransactionID string `json:"transaction_id"`
}

var _ bodyer = ReqFundsGetfundsflowlist{}

func (x ReqFundsGetfundsflowlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsGetfundsflowlist struct {
	CommonResp
	FlowIds []string `json:"flow_ids"`
	HasMore bool     `json:"has_more"`
	NextKey string   `json:"next_key"`
}

var _ bodyer = RespFundsGetfundsflowlist{}

func (x RespFundsGetfundsflowlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsGetfundsflowlist(req ReqFundsGetfundsflowlist) (RespFundsGetfundsflowlist, error) {
	var resp RespFundsGetfundsflowlist
	err := c.executeWXApiPost("/channels/ec/funds/getfundsflowlist", req, &resp, true)
	if err != nil {
		return RespFundsGetfundsflowlist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsGetfundsflowlist{}, bizErr
	}
	return resp, nil
}
