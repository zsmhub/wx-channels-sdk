package apis

import (
	"encoding/json"
)

// 获取账户余额
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/getbalance.html

type ReqFundsGetbalance struct{}

var _ bodyer = ReqFundsGetbalance{}

func (x ReqFundsGetbalance) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsGetbalance struct {
	CommonResp
	AvailableAmount int    `json:"available_amount"`
	PendingAmount   int    `json:"pending_amount"`
	SubMchid        string `json:"sub_mchid"`
}

var _ bodyer = RespFundsGetbalance{}

func (x RespFundsGetbalance) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsGetbalance(req ReqFundsGetbalance) (RespFundsGetbalance, error) {
	var resp RespFundsGetbalance
	err := c.executeWXApiPost("/channels/ec/funds/getbalance", req, &resp, true)
	if err != nil {
		return RespFundsGetbalance{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsGetbalance{}, bizErr
	}
	return resp, nil
}
