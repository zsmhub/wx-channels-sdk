package apis

import (
	"encoding/json"
)

// 获取地址详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/address/get.html

type ReqMerchantAddressGet struct {
	AddressID string `json:"address_id"`
}

var _ bodyer = ReqMerchantAddressGet{}

func (x ReqMerchantAddressGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantAddressGet struct {
	AddressDetail struct {
		AddressID   int `json:"address_id"`
		AddressInfo struct {
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
		} `json:"address_info"`
		AddressType struct {
			Pickup   int `json:"pickup"`
			SameCity int `json:"same_city"`
		} `json:"address_type"`
		CreateTime  int    `json:"create_time"`
		DefaultRecv int    `json:"default_recv"`
		DefaultSend int    `json:"default_send"`
		Landline    string `json:"landline"`
		Name        string `json:"name"`
		RecvAddr    int    `json:"recv_addr"`
		Remark      string `json:"remark"`
		SendAddr    int    `json:"send_addr"`
		UpdateTime  int    `json:"update_time"`
	} `json:"address_detail"`
	CommonResp
}

var _ bodyer = RespMerchantAddressGet{}

func (x RespMerchantAddressGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantAddressGet(req ReqMerchantAddressGet) (RespMerchantAddressGet, error) {
	var resp RespMerchantAddressGet
	err := c.executeWXApiPost("/channels/ec/merchant/address/get", req, &resp, true)
	if err != nil {
		return RespMerchantAddressGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantAddressGet{}, bizErr
	}
	return resp, nil
}
