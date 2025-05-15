package apis

import (
	"encoding/json"
)

// 下架商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/delisting.html

type ReqProductDelisting struct {
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqProductDelisting{}

func (x ReqProductDelisting) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductDelisting struct {
	CommonResp
}

var _ bodyer = RespProductDelisting{}

func (x RespProductDelisting) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductDelisting(req ReqProductDelisting) (RespProductDelisting, error) {
	var resp RespProductDelisting
	err := c.executeWXApiPost("/channels/ec/product/delisting", req, &resp, true)
	if err != nil {
		return RespProductDelisting{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductDelisting{}, bizErr
	}
	return resp, nil
}
