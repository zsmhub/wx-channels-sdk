package apis

import (
	"encoding/json"
)

// 修改订单备注
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/merchantnotes_update.html

type ReqOrderMerchantnotesUpdate struct {
	OrderId       string `json:"order_id"`
	MerchantNotes string `json:"merchant_notes"`
}

var _ bodyer = ReqOrderMerchantnotesUpdate{}

func (x ReqOrderMerchantnotesUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderMerchantnotesUpdate struct {
	CommonResp
}

var _ bodyer = RespOrderMerchantnotesUpdate{}

func (x RespOrderMerchantnotesUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderMerchantnotesUpdate(req ReqOrderMerchantnotesUpdate) (RespOrderMerchantnotesUpdate, error) {
	var resp RespOrderMerchantnotesUpdate
	err := c.executeWXApiPost("/channels/ec/order/merchantnotes/update", req, &resp, true)
	if err != nil {
		return RespOrderMerchantnotesUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderMerchantnotesUpdate{}, bizErr
	}
	return resp, nil
}
