package apis

import (
	"encoding/json"
)

// 更新区域仓库存
// 文档：https://developers.weixin.qq.com/doc/channels/API/warehouse/update_stock.html

type ReqWarehouseStockUpdate struct {
	DiffType       int    `json:"diff_type"`
	Num            int    `json:"num"`
	OutWarehouseID string `json:"out_warehouse_id"`
	ProductID      string `json:"product_id"`
	SkuID          string `json:"sku_id"`
}

var _ bodyer = ReqWarehouseStockUpdate{}

func (x ReqWarehouseStockUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWarehouseStockUpdate struct {
	CommonResp
}

var _ bodyer = RespWarehouseStockUpdate{}

func (x RespWarehouseStockUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWarehouseStockUpdate(req ReqWarehouseStockUpdate) (RespWarehouseStockUpdate, error) {
	var resp RespWarehouseStockUpdate
	err := c.executeWXApiPost("/channels/ec/warehouse/stock/update", req, &resp, true)
	if err != nil {
		return RespWarehouseStockUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWarehouseStockUpdate{}, bizErr
	}
	return resp, nil
}
