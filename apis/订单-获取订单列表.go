package apis

import (
	"encoding/json"
)

// 获取订单列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/list_get.html

type ReqOrderListGet struct {
	CreateTimeRange TimeRange `json:"create_time_range"`
	UpdateTimeRange TimeRange `json:"update_time_range"`
	Status          int       `json:"status"`
	Openid          string    `json:"openid"`
	NextKey         string    `json:"next_key"`
	PageSize        int       `json:"page_size"`
}

type TimeRange struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

var _ bodyer = ReqOrderListGet{}

func (x ReqOrderListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderListGet struct {
	CommonResp
	HasMore     bool     `json:"has_more"`
	NextKey     string   `json:"next_key"`
	OrderIDList []string `json:"order_id_list"`
}

var _ bodyer = RespOrderListGet{}

func (x RespOrderListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderListGet(req ReqOrderListGet) (RespOrderListGet, error) {
	var resp RespOrderListGet
	err := c.executeWXApiPost("/channels/ec/order/list/get", req, &resp, true)
	if err != nil {
		return RespOrderListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderListGet{}, bizErr
	}
	return resp, nil
}
