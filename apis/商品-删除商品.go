package apis

import (
	"encoding/json"
)

// 删除商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/delete.html

type ReqProductDelete struct {
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqProductDelete{}

func (x ReqProductDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductDelete struct {
	CommonResp
}

var _ bodyer = RespProductDelete{}

func (x RespProductDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductDelete(req ReqProductDelete) (RespProductDelete, error) {
	var resp RespProductDelete
	err := c.executeWXApiPost("/channels/ec/product/delete", req, &resp, true)
	if err != nil {
		return RespProductDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductDelete{}, bizErr
	}
	return resp, nil
}
