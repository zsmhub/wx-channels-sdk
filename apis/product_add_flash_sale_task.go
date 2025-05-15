package apis

import (
	"encoding/json"
)

// 添加限时抢购任务
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/limiteddiscounttask/add.html

type ReqProductLimiteddiscounttaskAdd struct {
	EndTime             int `json:"end_time"`
	LimitedDiscountSkus []struct {
		SalePrice int    `json:"sale_price"`
		SaleStock int    `json:"sale_stock"`
		SkuID     string `json:"sku_id"`
	} `json:"limited_discount_skus"`
	ProductID string `json:"product_id"`
	StartTime int    `json:"start_time"`
}

var _ bodyer = ReqProductLimiteddiscounttaskAdd{}

func (x ReqProductLimiteddiscounttaskAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductLimiteddiscounttaskAdd struct {
	CommonResp
	TaskID string `json:"task_id"`
}

var _ bodyer = RespProductLimiteddiscounttaskAdd{}

func (x RespProductLimiteddiscounttaskAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductLimiteddiscounttaskAdd(req ReqProductLimiteddiscounttaskAdd) (RespProductLimiteddiscounttaskAdd, error) {
	var resp RespProductLimiteddiscounttaskAdd
	err := c.executeWXApiPost("/channels/ec/product/limiteddiscounttask/add", req, &resp, true)
	if err != nil {
		return RespProductLimiteddiscounttaskAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductLimiteddiscounttaskAdd{}, bizErr
	}
	return resp, nil
}
