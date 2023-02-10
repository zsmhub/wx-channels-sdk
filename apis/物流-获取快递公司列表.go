package apis

import (
	"encoding/json"
)

// 获取快递公司列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/deliverycompanylist_get.html

type ReqOrderDeliverycompanylistGet struct{}

var _ bodyer = ReqOrderDeliverycompanylistGet{}

func (x ReqOrderDeliverycompanylistGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderDeliverycompanylistGet struct {
	CompanyList []struct {
		DeliveryID   string `json:"delivery_id"`
		DeliveryName string `json:"delivery_name"`
	} `json:"company_list"`
	CommonResp
}

var _ bodyer = RespOrderDeliverycompanylistGet{}

func (x RespOrderDeliverycompanylistGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderDeliverycompanylistGet(req ReqOrderDeliverycompanylistGet) (RespOrderDeliverycompanylistGet, error) {
	var resp RespOrderDeliverycompanylistGet
	err := c.executeWXApiPost("/channels/ec/order/deliverycompanylist/get", req, &resp, true)
	if err != nil {
		return RespOrderDeliverycompanylistGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderDeliverycompanylistGet{}, bizErr
	}
	return resp, nil
}
