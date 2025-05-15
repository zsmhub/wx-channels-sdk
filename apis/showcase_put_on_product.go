package apis

import (
	"encoding/json"
)

// 上架商品到橱窗
// 文档：https://developers.weixin.qq.com/doc/channels/API/windowproduct/add.html

type ReqWindowProductAdd struct {
	Appid           string `json:"appid"`
	IsHideForWindow bool   `json:"is_hide_for_window"`
	ProductID       string `json:"product_id"`
}

var _ bodyer = ReqWindowProductAdd{}

func (x ReqWindowProductAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWindowProductAdd struct {
	CommonResp
}

var _ bodyer = RespWindowProductAdd{}

func (x RespWindowProductAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWindowProductAdd(req ReqWindowProductAdd) (RespWindowProductAdd, error) {
	var resp RespWindowProductAdd
	err := c.executeWXApiPost("/channels/ec/window/product/add", req, &resp, true)
	if err != nil {
		return RespWindowProductAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWindowProductAdd{}, bizErr
	}
	return resp, nil
}
