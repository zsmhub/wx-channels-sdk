package apis

import (
	"encoding/json"
)

// 删除地址
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/address/delete.html

type ReqMerchantAddressDelete struct {
	AddressID string `json:"address_id"`
}

var _ bodyer = ReqMerchantAddressDelete{}

func (x ReqMerchantAddressDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantAddressDelete struct {
	CommonResp
}

var _ bodyer = RespMerchantAddressDelete{}

func (x RespMerchantAddressDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantAddressDelete(req ReqMerchantAddressDelete) (RespMerchantAddressDelete, error) {
	var resp RespMerchantAddressDelete
	err := c.executeWXApiPost("/channels/ec/merchant/address/delete", req, &resp, true)
	if err != nil {
		return RespMerchantAddressDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantAddressDelete{}, bizErr
	}
	return resp, nil
}
