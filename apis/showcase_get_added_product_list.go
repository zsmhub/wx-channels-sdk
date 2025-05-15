package apis

import (
	"encoding/json"
)

// 获取已添加到橱窗的商品列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/windowproduct/list_get.html

type ReqWindowProductListGet struct {
	Appid        string `json:"appid"`
	NeedTotalNum bool   `json:"need_total_num"`
	PageIndex    int    `json:"page_index"`
	PageSize     int    `json:"page_size"`
}

var _ bodyer = ReqWindowProductListGet{}

func (x ReqWindowProductListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWindowProductListGet struct {
	CommonResp
	LastBuffer string `json:"last_buffer"`
	Products   []struct {
		Appid     string `json:"appid"`
		ProductID string `json:"product_id"`
	} `json:"products"`
	TotalNum int `json:"total_num"`
}

var _ bodyer = RespWindowProductListGet{}

func (x RespWindowProductListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWindowProductListGet(req ReqWindowProductListGet) (RespWindowProductListGet, error) {
	var resp RespWindowProductListGet
	err := c.executeWXApiPost("/channels/ec/window/product/list/get", req, &resp, true)
	if err != nil {
		return RespWindowProductListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWindowProductListGet{}, bizErr
	}
	return resp, nil
}
