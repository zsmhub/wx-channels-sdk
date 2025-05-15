package apis

import (
	"encoding/json"
)

// 更新运费模版
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/updatefreighttemplate.html

type ReqMerchantUpdatefreighttemplate struct {
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
		IsDefault      int    `json:"is_default"`
		Name           string `json:"name"`
		SendTime       string `json:"send_time"`
		ShippingMethod string `json:"shipping_method"`
		ValuationType  string `json:"valuation_type"`
	} `json:"freight_template"`
}

var _ bodyer = ReqMerchantUpdatefreighttemplate{}

func (x ReqMerchantUpdatefreighttemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantUpdatefreighttemplate struct {
	CommonResp
	TemplateID string `json:"template_id"`
}

var _ bodyer = RespMerchantUpdatefreighttemplate{}

func (x RespMerchantUpdatefreighttemplate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantUpdatefreighttemplate(req ReqMerchantUpdatefreighttemplate) (RespMerchantUpdatefreighttemplate, error) {
	var resp RespMerchantUpdatefreighttemplate
	err := c.executeWXApiPost("/channels/ec/merchant/updatefreighttemplate", req, &resp, true)
	if err != nil {
		return RespMerchantUpdatefreighttemplate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantUpdatefreighttemplate{}, bizErr
	}
	return resp, nil
}
