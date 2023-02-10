package apis

import (
	"encoding/json"
)

// 快速更新库存
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/stock_update.html

type ReqProductStockUpdate struct {
	DiffType  int    `json:"diff_type"`
	Num       int    `json:"num"`
	ProductID string `json:"product_id"`
	SkuID     string `json:"sku_id"`
}

var _ bodyer = ReqProductStockUpdate{}

func (x ReqProductStockUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductStockUpdate struct {
	CommonResp
}

var _ bodyer = RespProductStockUpdate{}

func (x RespProductStockUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductStockUpdate(req ReqProductStockUpdate) (RespProductStockUpdate, error) {
	var resp RespProductStockUpdate
	err := c.executeWXApiPost("/channels/ec/product/stock/update", req, &resp, true)
	if err != nil {
		return RespProductStockUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductStockUpdate{}, bizErr
	}
	return resp, nil
}
