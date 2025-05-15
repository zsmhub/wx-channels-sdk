package apis

import (
	"encoding/json"
)

// 获取售后单列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/aftersale/getaftersalelist.html

type ReqAftersaleGetaftersalelist struct {
	BeginCreateTime int    `json:"begin_create_time"`
	EndCreateTime   int    `json:"end_create_time"`
	NextKey         string `json:"next_key"`
}

var _ bodyer = ReqAftersaleGetaftersalelist{}

func (x ReqAftersaleGetaftersalelist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleGetaftersalelist struct {
	AfterSaleOrderIDList []string `json:"after_sale_order_id_list"`
	CommonResp
	HasMore bool   `json:"has_more"`
	NextKey string `json:"next_key"`
}

var _ bodyer = RespAftersaleGetaftersalelist{}

func (x RespAftersaleGetaftersalelist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleGetaftersalelist(req ReqAftersaleGetaftersalelist) (RespAftersaleGetaftersalelist, error) {
	var resp RespAftersaleGetaftersalelist
	err := c.executeWXApiPost("/channels/ec/aftersale/getaftersalelist", req, &resp, true)
	if err != nil {
		return RespAftersaleGetaftersalelist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleGetaftersalelist{}, bizErr
	}
	return resp, nil
}
