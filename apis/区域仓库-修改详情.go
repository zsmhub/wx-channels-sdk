package apis

import (
	"encoding/json"
)

// 修改详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/update_detail.html

type ReqWarehouseDetailUpdate struct {
	Intro          string `json:"intro"`
	Name           string `json:"name"`
	OutWarehouseID string `json:"out_warehouse_id"`
}

var _ bodyer = ReqWarehouseDetailUpdate{}

func (x ReqWarehouseDetailUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseDetailUpdate struct {
	CommonResp
}

var _ bodyer = RespWarehouseDetailUpdate{}

func (x RespWarehouseDetailUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseDetailUpdate(req ReqWarehouseDetailUpdate) (RespWarehouseDetailUpdate, error) {
	var resp RespWarehouseDetailUpdate
	err := c.executeWXApiPost("/channels/ec/warehouse/detail/update", req, &resp, true)
	if err != nil {
		return RespWarehouseDetailUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseDetailUpdate{}, bizErr
	}
	return resp, nil
}
