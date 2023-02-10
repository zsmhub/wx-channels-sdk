package apis

import (
	"encoding/json"
)

// 获取实时库存
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/get_stock.html

type ReqProductStockGet struct {
	ProductId string `json:"product_id"`
	SkuId     string `json:"sku_id"`
}

var _ bodyer = ReqProductStockGet{}

func (x ReqProductStockGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductStockGet struct {
	Data struct {
		NormalStockNum          int `json:"normal_stock_num"`
		LimitedDiscountStockNum int `json:"limited_discount_stock_num"`
		WarehouseStocks         []struct {
			OutWarehouseId string `json:"out_warehouse_id"`
			Num            int    `json:"num"`
		} `json:"warehouse_stocks"`
		TotalStockNum int `json:"total_stock_num"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespProductStockGet{}

func (x RespProductStockGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductStockGet(req ReqProductStockGet) (RespProductStockGet, error) {
	var resp RespProductStockGet
	err := c.executeWXApiPost("/channels/ec/product/stock/get", req, &resp, true)
	if err != nil {
		return RespProductStockGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductStockGet{}, bizErr
	}
	return resp, nil
}
