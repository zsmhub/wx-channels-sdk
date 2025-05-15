package apis

import (
	"encoding/json"
)

// 订单搜索
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/search.html

type ReqOrderSearch struct {
	NextKey               string                        `json:"next_key"`
	OnAftersaleOrderExist *int                          `json:"on_aftersale_order_exist,omitempty"`
	PageSize              int                           `json:"page_size"`
	SearchCondition       ReqOrderSearchSearchCondition `json:"search_condition"`
	Status                int                           `json:"status,omitempty"`
}

type ReqOrderSearchSearchCondition struct {
	CustomerNotes string `json:"customer_notes,omitempty"`
	MerchantNotes string `json:"merchant_notes,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
	SkuCode       string `json:"sku_code,omitempty"`
	TelNumber     string `json:"tel_number,omitempty"`
	Title         string `json:"title,omitempty"`
	UserName      string `json:"user_name,omitempty"`
}

var _ bodyer = ReqOrderSearch{}

func (x ReqOrderSearch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderSearch struct {
	CommonResp
	HasMore bool     `json:"has_more"`
	NextKey string   `json:"next_key"`
	Orders  []string `json:"orders"`
}

var _ bodyer = RespOrderSearch{}

func (x RespOrderSearch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderSearch(req ReqOrderSearch) (RespOrderSearch, error) {
	var resp RespOrderSearch
	err := c.executeWXApiPost("/channels/ec/order/search", req, &resp, true)
	if err != nil {
		return RespOrderSearch{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderSearch{}, bizErr
	}
	return resp, nil
}
