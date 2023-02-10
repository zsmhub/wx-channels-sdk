package apis

import (
	"encoding/json"
)

// 增加运费模版
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/addfreighttemplate.html

type ReqMerchantAddfreighttemplate struct {
	FreightTemplate struct {
		AddressInfo struct {
			CityName     string `json:"city_name"`
			CountyName   string `json:"county_name"`
			DetailInfo   string `json:"detail_info"`
			NationalCode string `json:"national_code"`
			PostalCode   string `json:"postal_code"`
			ProvinceName string `json:"province_name"`
			TelNumber    string `json:"tel_number"`
			UserName     string `json:"user_name"`
		} `json:"address_info"`
		DeliveryType   string `json:"delivery_type"`
		IsDefault      bool   `json:"is_default"`
		Name           string `json:"name"`
		SendTime       string `json:"send_time"`
		ShippingMethod string `json:"shipping_method"`
		ValuationType  string `json:"valuation_type"`
	} `json:"freight_template"`
}

var _ bodyer = ReqMerchantAddfreighttemplate{}

func (x ReqMerchantAddfreighttemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantAddfreighttemplate struct {
	CommonResp
	TemplateID string `json:"template_id"`
}

var _ bodyer = RespMerchantAddfreighttemplate{}

func (x RespMerchantAddfreighttemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantAddfreighttemplate(req ReqMerchantAddfreighttemplate) (RespMerchantAddfreighttemplate, error) {
	var resp RespMerchantAddfreighttemplate
	err := c.executeWXApiPost("/channels/ec/merchant/addfreighttemplate", req, &resp, true)
	if err != nil {
		return RespMerchantAddfreighttemplate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantAddfreighttemplate{}, bizErr
	}
	return resp, nil
}
