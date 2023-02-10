package apis

import (
	"encoding/json"
)

// 获取区域仓库存数量
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/get_stock.html

type ReqWarehouseStockGet struct {
	OutWarehouseID string `json:"out_warehouse_id"`
	ProductID      string `json:"product_id"`
	SkuID          string `json:"sku_id"`
}

var _ bodyer = ReqWarehouseStockGet{}

func (x ReqWarehouseStockGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseStockGet struct {
	Data struct {
		Num int `json:"num"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespWarehouseStockGet{}

func (x RespWarehouseStockGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseStockGet(req ReqWarehouseStockGet) (RespWarehouseStockGet, error) {
	var resp RespWarehouseStockGet
	err := c.executeWXApiPost("/channels/ec/warehouse/stock/get", req, &resp, true)
	if err != nil {
		return RespWarehouseStockGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseStockGet{}, bizErr
	}
	return resp, nil
}
