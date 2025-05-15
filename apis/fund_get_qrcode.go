package apis

import (
	"encoding/json"
)

// 获取二维码
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/qrcode/get.html

type ReqQrcodeGet struct {
	QrcodeTicket string `json:"qrcode_ticket"`
}

var _ bodyer = ReqQrcodeGet{}

func (x ReqQrcodeGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespQrcodeGet struct {
	CommonResp
	QrcodeBuf string `json:"qrcode_buf"`
}

var _ bodyer = RespQrcodeGet{}

func (x RespQrcodeGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecQrcodeGet(req ReqQrcodeGet) (RespQrcodeGet, error) {
	var resp RespQrcodeGet
	err := c.executeWXApiPost("/shop/funds/qrcode/get", req, &resp, true)
	if err != nil {
		return RespQrcodeGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespQrcodeGet{}, bizErr
	}
	return resp, nil
}
