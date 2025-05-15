package apis

import (
	"encoding/json"
)

// 订单地址
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/address_update.html

type ReqOrderAddressUpdate struct {
	OrderID     string      `json:"order_id"`
	UserAddress UserAddress `json:"user_address"`
}

type UserAddress struct {
	CityName     string `json:"city_name"`
	CountyName   string `json:"county_name"`
	DetailInfo   string `json:"detail_info"`
	NationalCode string `json:"national_code"`
	PostalCode   string `json:"postal_code"`
	ProvinceName string `json:"province_name"`
	TelNumber    string `json:"tel_number"`
	UserName     string `json:"user_name"`
}

var _ bodyer = ReqOrderAddressUpdate{}

func (x ReqOrderAddressUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderAddressUpdate struct {
	CommonResp
}

var _ bodyer = RespOrderAddressUpdate{}

func (x RespOrderAddressUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderAddressUpdate(req ReqOrderAddressUpdate) (RespOrderAddressUpdate, error) {
	var resp RespOrderAddressUpdate
	err := c.executeWXApiPost("/channels/ec/order/address/update", req, &resp, true)
	if err != nil {
		return RespOrderAddressUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderAddressUpdate{}, bizErr
	}
	return resp, nil
}
