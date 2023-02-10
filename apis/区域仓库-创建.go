package apis

import (
	"encoding/json"
)

// 创建
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/create.html

type ReqWarehouseCreate struct {
	CoverLocations []struct {
		AddressID1 int `json:"address_id1"`
		AddressID2 int `json:"address_id2"`
		AddressID3 int `json:"address_id3"`
		AddressID4 int `json:"address_id4"`
	} `json:"cover_locations"`
	Intro          string `json:"intro"`
	Name           string `json:"name"`
	OutWarehouseID string `json:"out_warehouse_id"`
}

var _ bodyer = ReqWarehouseCreate{}

func (x ReqWarehouseCreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseCreate struct {
	CommonResp
}

var _ bodyer = RespWarehouseCreate{}

func (x RespWarehouseCreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseCreate(req ReqWarehouseCreate) (RespWarehouseCreate, error) {
	var resp RespWarehouseCreate
	err := c.executeWXApiPost("/channels/ec/warehouse/create", req, &resp, true)
	if err != nil {
		return RespWarehouseCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseCreate{}, bizErr
	}
	return resp, nil
}
