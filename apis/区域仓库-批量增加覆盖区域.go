package apis

import (
	"encoding/json"
)

// 批量增加覆盖区域
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/add_coverlocations.html

type ReqWarehouseCoverlocationsAdd struct {
	CoverLocations []struct {
		AddressID1 int `json:"address_id1"`
		AddressID2 int `json:"address_id2"`
		AddressID3 int `json:"address_id3"`
		AddressID4 int `json:"address_id4"`
	} `json:"cover_locations"`
	OutWarehouseID string `json:"out_warehouse_id"`
}

var _ bodyer = ReqWarehouseCoverlocationsAdd{}

func (x ReqWarehouseCoverlocationsAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseCoverlocationsAdd struct {
	CommonResp
}

var _ bodyer = RespWarehouseCoverlocationsAdd{}

func (x RespWarehouseCoverlocationsAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseCoverlocationsAdd(req ReqWarehouseCoverlocationsAdd) (RespWarehouseCoverlocationsAdd, error) {
	var resp RespWarehouseCoverlocationsAdd
	err := c.executeWXApiPost("/channels/ec/warehouse/coverlocations/add", req, &resp, true)
	if err != nil {
		return RespWarehouseCoverlocationsAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseCoverlocationsAdd{}, bizErr
	}
	return resp, nil
}
