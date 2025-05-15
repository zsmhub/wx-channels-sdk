package apis

import (
	"encoding/json"
)

// 拉取限时抢购任务列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/limiteddiscounttask/list_get.html

type ReqProductLimiteddiscounttaskListGet struct {
	NextKey  string `json:"next_key"`
	PageSize int    `json:"page_size"`
	Status   int    `json:"status"`
}

var _ bodyer = ReqProductLimiteddiscounttaskListGet{}

func (x ReqProductLimiteddiscounttaskListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductLimiteddiscounttaskListGet struct {
	CommonResp
	LimitedDiscountTasks []struct {
		CreateTime          int `json:"create_time"`
		EndTime             int `json:"end_time"`
		LimitedDiscountSkus []struct {
			SalePrice int    `json:"sale_price"`
			SaleStock int    `json:"sale_stock"`
			SkuID     string `json:"sku_id"`
		} `json:"limited_discount_skus"`
		ProductID string `json:"product_id"`
		StartTime int    `json:"start_time"`
		Status    int    `json:"status"`
		TaskID    string `json:"task_id"`
	} `json:"limited_discount_tasks"`
	NextKey  string `json:"next_key"`
	TotalNum int    `json:"total_num"`
}

var _ bodyer = RespProductLimiteddiscounttaskListGet{}

func (x RespProductLimiteddiscounttaskListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductLimiteddiscounttaskListGet(req ReqProductLimiteddiscounttaskListGet) (RespProductLimiteddiscounttaskListGet, error) {
	var resp RespProductLimiteddiscounttaskListGet
	err := c.executeWXApiPost("/channels/ec/product/limiteddiscounttask/list/get", req, &resp, true)
	if err != nil {
		return RespProductLimiteddiscounttaskListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductLimiteddiscounttaskListGet{}, bizErr
	}
	return resp, nil
}
