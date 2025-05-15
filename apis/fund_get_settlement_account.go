package apis

import (
	"encoding/json"
)

// 获取结算账户
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/getbankacct.html

type ReqFundsGetbankacct struct{}

var _ bodyer = ReqFundsGetbankacct{}

func (x ReqFundsGetbankacct) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsGetbankacct struct {
	CommonResp
	AccountInfo struct {
		AccountBank     string `json:"account_bank"`
		AccountNumber   string `json:"account_number"`
		BankAccountType string `json:"bank_account_type"`
		BankAddressCode string `json:"bank_address_code"`
	} `json:"account_info"`
}

var _ bodyer = RespFundsGetbankacct{}

func (x RespFundsGetbankacct) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsGetbankacct(req ReqFundsGetbankacct) (RespFundsGetbankacct, error) {
	var resp RespFundsGetbankacct
	err := c.executeWXApiPost("/channels/ec/funds/getbankacct", req, &resp, true)
	if err != nil {
		return RespFundsGetbankacct{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsGetbankacct{}, bizErr
	}
	return resp, nil
}
