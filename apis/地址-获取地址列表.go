package apis

import (
	"encoding/json"
)

// 添加地址
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/address/add.html

type ReqMerchantAddressList struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

var _ bodyer = ReqMerchantAddressList{}

func (x ReqMerchantAddressList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantAddressList struct {
	CommonResp
	AddressIdList []string `json:"address_id_list"`
}

var _ bodyer = RespMerchantAddressList{}

func (x RespMerchantAddressList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantAddressList(req ReqMerchantAddressList) (RespMerchantAddressList, error) {
	var resp RespMerchantAddressList
	err := c.executeWXApiPost("/channels/ec/merchant/address/list", req, &resp, true)
	if err != nil {
		return RespMerchantAddressList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantAddressList{}, bizErr
	}
	return resp, nil
}
