package apis

import (
	"encoding/json"
)

// 查询扫码状态
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/qrcode/check.html

type ReqQrcodeCheck struct {
	QrcodeTicket string `json:"qrcode_ticket"`
}

var _ bodyer = ReqQrcodeCheck{}

func (x ReqQrcodeCheck) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespQrcodeCheck struct {
	CommonResp
	ScanUserType     int    `json:"scan_user_type"`
	SelfCheckErrCode int    `json:"self_check_err_code"`
	SelfCheckErrMsg  string `json:"self_check_err_msg"`
	Status           int    `json:"status"`
}

var _ bodyer = RespQrcodeCheck{}

func (x RespQrcodeCheck) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecQrcodeCheck(req ReqQrcodeCheck) (RespQrcodeCheck, error) {
	var resp RespQrcodeCheck
	err := c.executeWXApiPost("/shop/funds/qrcode/check", req, &resp, true)
	if err != nil {
		return RespQrcodeCheck{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespQrcodeCheck{}, bizErr
	}
	return resp, nil
}
