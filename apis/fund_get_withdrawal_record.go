package apis

import (
	"encoding/json"
)

// 获取提现记录
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/getwithdrawdetail.html

type ReqFundsGetwithdrawdetail struct {
	WithdrawID string `json:"withdraw_id"`
}

var _ bodyer = ReqFundsGetwithdrawdetail{}

func (x ReqFundsGetwithdrawdetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsGetwithdrawdetail struct {
	CommonResp
	Amount     int    `json:"amount"`
	BankMemo   string `json:"bank_memo"`
	BankName   string `json:"bank_name"`
	BankNum    string `json:"bank_num"`
	CreateTime int    `json:"create_time"`
	Reason     string `json:"reason"`
	Remark     string `json:"remark"`
	Status     string `json:"status"`
	UpdateTime int    `json:"update_time"`
}

var _ bodyer = RespFundsGetwithdrawdetail{}

func (x RespFundsGetwithdrawdetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsGetwithdrawdetail(req ReqFundsGetwithdrawdetail) (RespFundsGetwithdrawdetail, error) {
	var resp RespFundsGetwithdrawdetail
	err := c.executeWXApiPost("/channels/ec/funds/getwithdrawdetail", req, &resp, true)
	if err != nil {
		return RespFundsGetwithdrawdetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsGetwithdrawdetail{}, bizErr
	}
	return resp, nil
}
