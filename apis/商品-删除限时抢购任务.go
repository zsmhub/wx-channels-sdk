package apis

import (
	"encoding/json"
)

// 删除限时抢购任务
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/limiteddiscounttask/delete.html

type ReqProductLimiteddiscounttaskDelete struct {
	TaskID string `json:"task_id"`
}

var _ bodyer = ReqProductLimiteddiscounttaskDelete{}

func (x ReqProductLimiteddiscounttaskDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductLimiteddiscounttaskDelete struct {
	CommonResp
}

var _ bodyer = RespProductLimiteddiscounttaskDelete{}

func (x RespProductLimiteddiscounttaskDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductLimiteddiscounttaskDelete(req ReqProductLimiteddiscounttaskDelete) (RespProductLimiteddiscounttaskDelete, error) {
	var resp RespProductLimiteddiscounttaskDelete
	err := c.executeWXApiPost("/channels/ec/product/limiteddiscounttask/delete", req, &resp, true)
	if err != nil {
		return RespProductLimiteddiscounttaskDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductLimiteddiscounttaskDelete{}, bizErr
	}
	return resp, nil
}
