package apis

import (
	"encoding/json"
)

// 根据卡号查银行信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/bank/getbankbynum.html

type ReqGetbankbynum struct {
	AccountNumber string `json:"account_number"`
}

var _ bodyer = ReqGetbankbynum{}

func (x ReqGetbankbynum) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetbankbynum struct {
	CommonResp
	Data []struct {
		AccountBank string `json:"account_bank"`
		BankCode    string `json:"bank_code"`
		BankID      string `json:"bank_id"`
		BankName    string `json:"bank_name"`
		NeedBranch  bool   `json:"need_branch"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

var _ bodyer = RespGetbankbynum{}

func (x RespGetbankbynum) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetbankbynum(req ReqGetbankbynum) (RespGetbankbynum, error) {
	var resp RespGetbankbynum
	err := c.executeWXApiPost("/shop/funds/getbankbynum", req, &resp, true)
	if err != nil {
		return RespGetbankbynum{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetbankbynum{}, bizErr
	}
	return resp, nil
}
