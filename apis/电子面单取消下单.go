package apis

import (
	"encoding/json"
)

// 电子面单取消下单
// 文档：https://developers.weixin.qq.com/doc/channels/API/ewaybill/cancel_order.html

type ReqLogisticsEwaybillBizOrderCancel struct {
	DeliveryID      string `json:"delivery_id"`
	EwaybillOrderID string `json:"ewaybill_order_id"`
	WaybillID       string `json:"waybill_id"`
}

var _ bodyer = ReqLogisticsEwaybillBizOrderCancel{}

func (x ReqLogisticsEwaybillBizOrderCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespLogisticsEwaybillBizOrderCancel struct {
	DeliveryErrorMsg string `json:"delivery_error_msg"`
	CommonResp
}

var _ bodyer = RespLogisticsEwaybillBizOrderCancel{}

func (x RespLogisticsEwaybillBizOrderCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLogisticsEwaybillBizOrderCancel(req ReqLogisticsEwaybillBizOrderCancel) (RespLogisticsEwaybillBizOrderCancel, error) {
	var resp RespLogisticsEwaybillBizOrderCancel
	err := c.executeWXApiPost("/channels/ec/logistics/ewaybill/biz/order/cancel", req, &resp, true)
	if err != nil {
		return RespLogisticsEwaybillBizOrderCancel{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLogisticsEwaybillBizOrderCancel{}, bizErr
	}
	return resp, nil
}
