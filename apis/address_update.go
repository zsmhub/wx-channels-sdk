package apis

import (
	"encoding/json"
)

// 更新地址
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/address/update.html

type ReqMerchantAddressUpdate struct {
	AddressDetail MerchantAddressDetail `json:"address_detail"`
}

var _ bodyer = ReqMerchantAddressUpdate{}

func (x ReqMerchantAddressUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantAddressUpdate struct {
	CommonResp
}

var _ bodyer = RespMerchantAddressUpdate{}

func (x RespMerchantAddressUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantAddressUpdate(req ReqMerchantAddressUpdate) (RespMerchantAddressUpdate, error) {
	var resp RespMerchantAddressUpdate
	err := c.executeWXApiPost("/channels/ec/merchant/address/update", req, &resp, true)
	if err != nil {
		return RespMerchantAddressUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantAddressUpdate{}, bizErr
	}
	return resp, nil
}
