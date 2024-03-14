package apis

import (
	"encoding/json"
)

// 查询开通的快递公司列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/ewaybill/get_acctdeliverylist.html

type ReqLogisticsEwaybillBizDeliveryGet struct {
	Status int `json:"status"`
}

var _ bodyer = ReqLogisticsEwaybillBizDeliveryGet{}

func (x ReqLogisticsEwaybillBizDeliveryGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespLogisticsEwaybillBizDeliveryGet struct {
	CommonResp
	List []struct {
		DeliveryID   string `json:"delivery_id"`
		DeliveryName string `json:"delivery_name"`
	} `json:"list"`
	ShopID string `json:"shop_id"`
}

var _ bodyer = RespLogisticsEwaybillBizDeliveryGet{}

func (x RespLogisticsEwaybillBizDeliveryGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLogisticsEwaybillBizDeliveryGet(req ReqLogisticsEwaybillBizDeliveryGet) (RespLogisticsEwaybillBizDeliveryGet, error) {
	var resp RespLogisticsEwaybillBizDeliveryGet
	err := c.executeWXApiPost("/channels/ec/logistics/ewaybill/biz/delivery/get", req, &resp, true)
	if err != nil {
		return RespLogisticsEwaybillBizDeliveryGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLogisticsEwaybillBizDeliveryGet{}, bizErr
	}
	return resp, nil
}
