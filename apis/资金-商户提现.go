package apis

import (
	"encoding/json"
)

// 商户提现
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/submitwithdraw.html

type ReqFundsSubmitwithdraw struct {
	Amount   int    `json:"amount"`
	BankMemo string `json:"bank_memo"`
	Remark   string `json:"remark"`
}

var _ bodyer = ReqFundsSubmitwithdraw{}

func (x ReqFundsSubmitwithdraw) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsSubmitwithdraw struct {
	CommonResp
	QrcodeTicket string `json:"qrcode_ticket"`
}

var _ bodyer = RespFundsSubmitwithdraw{}

func (x RespFundsSubmitwithdraw) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsSubmitwithdraw(req ReqFundsSubmitwithdraw) (RespFundsSubmitwithdraw, error) {
	var resp RespFundsSubmitwithdraw
	err := c.executeWXApiPost("/channels/ec/funds/submitwithdraw", req, &resp, true)
	if err != nil {
		return RespFundsSubmitwithdraw{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsSubmitwithdraw{}, bizErr
	}
	return resp, nil
}
