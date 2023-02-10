package apis

import (
	"encoding/json"
)

// 设置指定地址下的仓的优先级
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/set_prioritysort.html

type ReqWarehouseAddressPrioritysortSet struct {
	AddressID1   int      `json:"address_id1"`
	AddressID2   int      `json:"address_id2"`
	AddressID3   int      `json:"address_id3"`
	AddressID4   int      `json:"address_id4"`
	PrioritySort []string `json:"priority_sort"`
}

var _ bodyer = ReqWarehouseAddressPrioritysortSet{}

func (x ReqWarehouseAddressPrioritysortSet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseAddressPrioritysortSet struct {
	CommonResp
}

var _ bodyer = RespWarehouseAddressPrioritysortSet{}

func (x RespWarehouseAddressPrioritysortSet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseAddressPrioritysortSet(req ReqWarehouseAddressPrioritysortSet) (RespWarehouseAddressPrioritysortSet, error) {
	var resp RespWarehouseAddressPrioritysortSet
	err := c.executeWXApiPost("/channels/ec/warehouse/address/prioritysort/set", req, &resp, true)
	if err != nil {
		return RespWarehouseAddressPrioritysortSet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseAddressPrioritysortSet{}, bizErr
	}
	return resp, nil
}
