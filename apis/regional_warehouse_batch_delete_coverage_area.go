package apis

import (
	"encoding/json"
)

// 批量删除覆盖区域
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/del_coverlocations.html

type ReqWarehouseCoverlocationsDel struct {
	CoverLocations []struct {
		AddressID1 int `json:"address_id1"`
		AddressID2 int `json:"address_id2"`
		AddressID3 int `json:"address_id3"`
		AddressID4 int `json:"address_id4"`
	} `json:"cover_locations"`
	OutWarehouseID string `json:"out_warehouse_id"`
}

var _ bodyer = ReqWarehouseCoverlocationsDel{}

func (x ReqWarehouseCoverlocationsDel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseCoverlocationsDel struct {
	CommonResp
}

var _ bodyer = RespWarehouseCoverlocationsDel{}

func (x RespWarehouseCoverlocationsDel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseCoverlocationsDel(req ReqWarehouseCoverlocationsDel) (RespWarehouseCoverlocationsDel, error) {
	var resp RespWarehouseCoverlocationsDel
	err := c.executeWXApiPost("/channels/ec/warehouse/coverlocations/del", req, &resp, true)
	if err != nil {
		return RespWarehouseCoverlocationsDel{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseCoverlocationsDel{}, bizErr
	}
	return resp, nil
}
