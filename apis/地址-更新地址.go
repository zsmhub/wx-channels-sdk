package apis

import (
	"encoding/json"
)

// 更新地址
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/address/update.html

type ReqMerchantAddressUpdate struct {
	AddressDetail struct {
		AddressID   string `json:"address_id"`
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
		DefaultRecv int    `json:"default_recv"`
		DefaultSend int    `json:"default_send"`
		Landline    string `json:"landline"`
		RecvAddr    int    `json:"recv_addr"`
		Remark      string `json:"remark"`
		SendAddr    int    `json:"send_addr"`
	} `json:"address_detail"`
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
