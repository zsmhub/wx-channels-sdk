package apis

import (
	"encoding/json"
)

// 下架橱窗商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/windowproduct/off.html

type ReqWindowProductOff struct {
	Appid     string `json:"appid"`
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqWindowProductOff{}

func (x ReqWindowProductOff) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWindowProductOff struct {
	CommonResp
}

var _ bodyer = RespWindowProductOff{}

func (x RespWindowProductOff) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWindowProductOff(req ReqWindowProductOff) (RespWindowProductOff, error) {
	var resp RespWindowProductOff
	err := c.executeWXApiPost("/channels/ec/window/product/off", req, &resp, true)
	if err != nil {
		return RespWindowProductOff{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWindowProductOff{}, bizErr
	}
	return resp, nil
}
