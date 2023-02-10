package apis

import (
	"encoding/json"
)

// 查询列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/get_list.html

type ReqWarehouseListGet struct {
	NextKey  string `json:"next_key"`
	PageSize int    `json:"page_size"`
}

var _ bodyer = ReqWarehouseListGet{}

func (x ReqWarehouseListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseListGet struct {
	Data struct {
		NextKey         string   `json:"next_key"`
		OutWarehouseIds []string `json:"out_warehouse_ids"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespWarehouseListGet{}

func (x RespWarehouseListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseListGet(req ReqWarehouseListGet) (RespWarehouseListGet, error) {
	var resp RespWarehouseListGet
	err := c.executeWXApiPost("/channels/ec/warehouse/list/get", req, &resp, true)
	if err != nil {
		return RespWarehouseListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseListGet{}, bizErr
	}
	return resp, nil
}
