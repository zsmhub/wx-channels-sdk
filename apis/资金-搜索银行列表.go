package apis

import (
	"encoding/json"
)

// 搜索银行列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/bank/getbanklist.html

type ReqGetbanklist struct {
	BankType int    `json:"bank_type"`
	KeyWords string `json:"key_words"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

var _ bodyer = ReqGetbanklist{}

func (x ReqGetbanklist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetbanklist struct {
	CommonResp
	Data []struct {
		AccountBank string `json:"account_bank"`
		BankCode    string `json:"bank_code"`
		BankID      string `json:"bank_id"`
		BankName    string `json:"bank_name"`
		BankType    int    `json:"bank_type"`
		NeedBranch  bool   `json:"need_branch"`
	} `json:"data"`
}

var _ bodyer = RespGetbanklist{}

func (x RespGetbanklist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetbanklist(req ReqGetbanklist) (RespGetbanklist, error) {
	var resp RespGetbanklist
	err := c.executeWXApiPost("/shop/funds/getbanklist", req, &resp, true)
	if err != nil {
		return RespGetbanklist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetbanklist{}, bizErr
	}
	return resp, nil
}
