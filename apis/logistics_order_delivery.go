package apis

import (
	"encoding/json"
)

// 订单发货
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/delivery_send.html

type ReqOrderDeliverySend struct {
	DeliveryList []DeliveryListItem `json:"delivery_list"`
	OrderID      string             `json:"order_id"`
}

type DeliveryListItem struct {
	DeliverType  int                   `json:"deliver_type"`
	DeliveryID   string                `json:"delivery_id"`
	ProductInfos []DeliveryProductInfo `json:"product_infos"`
	WaybillID    string                `json:"waybill_id"`
}
type DeliveryProductInfo struct {
	ProductCnt int    `json:"product_cnt"`
	ProductID  string `json:"product_id"`
	SkuID      string `json:"sku_id"`
}

var _ bodyer = ReqOrderDeliverySend{}

func (x ReqOrderDeliverySend) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderDeliverySend struct {
	CommonResp
}

var _ bodyer = RespOrderDeliverySend{}

func (x RespOrderDeliverySend) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderDeliverySend(req ReqOrderDeliverySend) (RespOrderDeliverySend, error) {
	var resp RespOrderDeliverySend
	err := c.executeWXApiPost("/channels/ec/order/delivery/send", req, &resp, true)
	if err != nil {
		return RespOrderDeliverySend{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderDeliverySend{}, bizErr
	}
	return resp, nil
}
