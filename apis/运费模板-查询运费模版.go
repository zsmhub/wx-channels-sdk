package apis

import (
	"encoding/json"
)

// 查询运费模版
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/getfreighttemplatedetail.html

type ReqMerchantGetfreighttemplatedetail struct {
	TemplateID string `json:"template_id"`
}

var _ bodyer = ReqMerchantGetfreighttemplatedetail{}

func (x ReqMerchantGetfreighttemplatedetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantGetfreighttemplatedetail struct {
	CommonResp
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
		AllConditionFreeDetail struct {
			ConditionFreeDetailList []interface{} `json:"condition_free_detail_list"`
		} `json:"all_condition_free_detail"`
		AllFreightCalcMethod struct {
			FreightCalcMethodList []interface{} `json:"freight_calc_method_list"`
		} `json:"all_freight_calc_method"`
		CreateTime   int           `json:"create_time"`
		DeliveryID   []interface{} `json:"delivery_id"`
		DeliveryType string        `json:"delivery_type"`
		IsDefault    int           `json:"is_default"`
		Name         string        `json:"name"`
		NotSendArea  struct {
			AddressInfo []interface{} `json:"address_info"`
		} `json:"not_send_area"`
		SendTime       string `json:"send_time"`
		ShippingMethod string `json:"shipping_method"`
		TemplateID     string `json:"template_id"`
		UpdateTime     int    `json:"update_time"`
		ValuationType  string `json:"valuation_type"`
	} `json:"freight_template"`
}

var _ bodyer = RespMerchantGetfreighttemplatedetail{}

func (x RespMerchantGetfreighttemplatedetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantGetfreighttemplatedetail(req ReqMerchantGetfreighttemplatedetail) (RespMerchantGetfreighttemplatedetail, error) {
	var resp RespMerchantGetfreighttemplatedetail
	err := c.executeWXApiPost("/channels/ec/merchant/getfreighttemplatedetail", req, &resp, true)
	if err != nil {
		return RespMerchantGetfreighttemplatedetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantGetfreighttemplatedetail{}, bizErr
	}
	return resp, nil
}
