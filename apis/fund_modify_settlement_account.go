package apis

import (
	"encoding/json"
)

// 修改结算账户
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/setbankacct.html

type ReqFundsSetbankacct struct {
	AccountInfo struct {
		AccountBank     string `json:"account_bank"`
		AccountNumber   string `json:"account_number"`
		BankAccountType string `json:"bank_account_type"`
		BankAddressCode string `json:"bank_address_code"`
	} `json:"account_info"`
}

var _ bodyer = ReqFundsSetbankacct{}

func (x ReqFundsSetbankacct) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsSetbankacct struct {
	CommonResp
}

var _ bodyer = RespFundsSetbankacct{}

func (x RespFundsSetbankacct) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsSetbankacct(req ReqFundsSetbankacct) (RespFundsSetbankacct, error) {
	var resp RespFundsSetbankacct
	err := c.executeWXApiPost("/channels/ec/funds/setbankacct", req, &resp, true)
	if err != nil {
		return RespFundsSetbankacct{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsSetbankacct{}, bizErr
	}
	return resp, nil
}
