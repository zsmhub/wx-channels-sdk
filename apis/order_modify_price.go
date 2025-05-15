package apis

import (
	"encoding/json"
)

// 修改订单价格
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/price_update.html

type ReqOrderPriceUpdate struct {
	ChangeExpress    bool                   `json:"change_express"`
	ChangeOrderInfos []ChangeOrderInfosItem `json:"change_order_infos"`
	ExpressFee       *int                   `json:"express_fee,omitempty"`
	OrderID          string                 `json:"order_id"`
}

type ChangeOrderInfosItem struct {
	ChangePrice int    `json:"change_price"`
	ProductID   string `json:"product_id"`
	SkuID       string `json:"sku_id"`
}

var _ bodyer = ReqOrderPriceUpdate{}

func (x ReqOrderPriceUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderPriceUpdate struct {
	CommonResp
}

var _ bodyer = RespOrderPriceUpdate{}

func (x RespOrderPriceUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderPriceUpdate(req ReqOrderPriceUpdate) (RespOrderPriceUpdate, error) {
	var resp RespOrderPriceUpdate
	err := c.executeWXApiPost("/channels/ec/order/price/update", req, &resp, true)
	if err != nil {
		return RespOrderPriceUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderPriceUpdate{}, bizErr
	}
	return resp, nil
}
