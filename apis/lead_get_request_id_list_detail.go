package apis

import (
	"encoding/json"
)

// 获取留资request_id列表详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/leads/get_leads_request_id.html

type ReqGetLeadsRequestId struct {
	LastBuffer       string `json:"last_buffer"`
	LeadsComponentID string `json:"leads_component_id"`
}

var _ bodyer = ReqGetLeadsRequestId{}

func (x ReqGetLeadsRequestId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetLeadsRequestId struct {
	ContinueFlag bool `json:"continue_flag"`
	CommonResp
	Item []struct {
		LiveDescription string `json:"live_description"`
		LiveStartTime   int    `json:"live_start_time"`
		RequestID       string `json:"request_id"`
	} `json:"item"`
	LastBuffer string `json:"last_buffer"`
}

var _ bodyer = RespGetLeadsRequestId{}

func (x RespGetLeadsRequestId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetLeadsRequestId(req ReqGetLeadsRequestId) (RespGetLeadsRequestId, error) {
	var resp RespGetLeadsRequestId
	err := c.executeWXApiPost("/channels/leads/get_leads_request_id", req, &resp, true)
	if err != nil {
		return RespGetLeadsRequestId{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetLeadsRequestId{}, bizErr
	}
	return resp, nil
}
