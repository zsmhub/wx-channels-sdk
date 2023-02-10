package apis

import (
	"encoding/json"
)

// 获取
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/get.html

type ReqWarehouseGet struct {
	OutWarehouseID string `json:"out_warehouse_id"`
}

var _ bodyer = ReqWarehouseGet{}

func (x ReqWarehouseGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseGet struct {
	Data struct {
		CoverLocations []struct {
			AddressID1 int `json:"address_id1"`
			AddressID2 int `json:"address_id2"`
			AddressID3 int `json:"address_id3"`
			AddressID4 int `json:"address_id4"`
		} `json:"cover_locations"`
		Intro          string `json:"intro"`
		Name           string `json:"name"`
		OutWarehouseID string `json:"out_warehouse_id"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespWarehouseGet{}

func (x RespWarehouseGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseGet(req ReqWarehouseGet) (RespWarehouseGet, error) {
	var resp RespWarehouseGet
	err := c.executeWXApiPost("/channels/ec/warehouse/get", req, &resp, true)
	if err != nil {
		return RespWarehouseGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseGet{}, bizErr
	}
	return resp, nil
}
