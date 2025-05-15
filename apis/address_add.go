package apis

import (
	"encoding/json"
)

// 添加地址
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/address/add.html

type ReqMerchantAddressAdd struct {
	AddressDetail MerchantAddressDetail `json:"address_detail"`
}

type MerchantAddressDetail struct {
	AddressID   string              `json:"address_id,omitempty"` // 新增地址不用传
	AddressInfo MerchantAddressInfo `json:"address_info"`
	AddressType MerchantAddressType `json:"address_type"`
	DefaultRecv bool                `json:"default_recv"`
	DefaultSend bool                `json:"default_send"`
	Landline    string              `json:"landline"`
	RecvAddr    bool                `json:"recv_addr"`
	Remark      string              `json:"remark"`
	SendAddr    bool                `json:"send_addr"`
}

type MerchantAddressInfo struct {
	CityName     string  `json:"city_name"`
	CountyName   string  `json:"county_name"`
	DetailInfo   string  `json:"detail_info"`
	HouseNumber  string  `json:"house_number"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	PostalCode   string  `json:"postal_code"`
	ProvinceName string  `json:"province_name"`
	TelNumber    string  `json:"tel_number"`
	UserName     string  `json:"user_name"`
}

type MerchantAddressType struct {
	Pickup   int `json:"pickup"`
	SameCity int `json:"same_city"`
}

var _ bodyer = ReqMerchantAddressAdd{}

func (x ReqMerchantAddressAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantAddressAdd struct {
	CommonResp
	AddressID string `json:"address_id"`
}

var _ bodyer = RespMerchantAddressAdd{}

func (x RespMerchantAddressAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantAddressAdd(req ReqMerchantAddressAdd) (RespMerchantAddressAdd, error) {
	var resp RespMerchantAddressAdd
	err := c.executeWXApiPost("/channels/ec/merchant/address/add", req, &resp, true)
	if err != nil {
		return RespMerchantAddressAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantAddressAdd{}, bizErr
	}
	return resp, nil
}
