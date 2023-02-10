package apis

import (
	"encoding/json"
)

// 获取指定地址下的仓的优先级
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/get_prioritysort.html

type ReqWarehouseAddressPrioritysortGet struct {
	AddressID1 int `json:"address_id1"`
	AddressID2 int `json:"address_id2"`
	AddressID3 int `json:"address_id3"`
	AddressID4 int `json:"address_id4"`
}

var _ bodyer = ReqWarehouseAddressPrioritysortGet{}

func (x ReqWarehouseAddressPrioritysortGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseAddressPrioritysortGet struct {
	Data struct {
		PrioritySort []string `json:"priority_sort"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespWarehouseAddressPrioritysortGet{}

func (x RespWarehouseAddressPrioritysortGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseAddressPrioritysortGet(req ReqWarehouseAddressPrioritysortGet) (RespWarehouseAddressPrioritysortGet, error) {
	var resp RespWarehouseAddressPrioritysortGet
	err := c.executeWXApiPost("/channels/ec/warehouse/address/prioritysort/get", req, &resp, true)
	if err != nil {
		return RespWarehouseAddressPrioritysortGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseAddressPrioritysortGet{}, bizErr
	}
	return resp, nil
}
