package apis

import (
	"encoding/json"
)

// 停止限时抢购任务
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/limiteddiscounttask/stop.html

type ReqProductLimiteddiscounttaskStop struct {
	TaskID string `json:"task_id"`
}

var _ bodyer = ReqProductLimiteddiscounttaskStop{}

func (x ReqProductLimiteddiscounttaskStop) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductLimiteddiscounttaskStop struct {
	CommonResp
}

var _ bodyer = RespProductLimiteddiscounttaskStop{}

func (x RespProductLimiteddiscounttaskStop) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductLimiteddiscounttaskStop(req ReqProductLimiteddiscounttaskStop) (RespProductLimiteddiscounttaskStop, error) {
	var resp RespProductLimiteddiscounttaskStop
	err := c.executeWXApiPost("/channels/ec/product/limiteddiscounttask/stop", req, &resp, true)
	if err != nil {
		return RespProductLimiteddiscounttaskStop{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductLimiteddiscounttaskStop{}, bizErr
	}
	return resp, nil
}
