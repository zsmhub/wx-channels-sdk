package apis

import (
	"encoding/json"
)

// 获取留资request_id列表详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/leads/get_leads_request_id.html

type ReqGetLeadsInfoByComponentId struct {
	LeadsComponentId string `json:"leads_component_id"`
	StartTime        int    `json:"start_time"`
	EndTime          int    `json:"end_time"`
	LastBuffer       string `json:"last_buffer"`
}

var _ bodyer = ReqGetLeadsInfoByComponentId{}

func (x ReqGetLeadsInfoByComponentId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetLeadsInfoByComponentId struct {
	CommonResp
	UserData []struct {
		Phone string `json:"phone"`
	} `json:"user_data"`
	LastBuffer   string `json:"last_buffer"`
	ContinueFlag bool   `json:"continue_flag"`
}

var _ bodyer = RespGetLeadsInfoByComponentId{}

func (x RespGetLeadsInfoByComponentId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetLeadsInfoByComponentId(req ReqGetLeadsInfoByComponentId) (RespGetLeadsInfoByComponentId, error) {
	var resp RespGetLeadsInfoByComponentId
	err := c.executeWXApiPost("/channels/leads/get_leads_info_by_component_id", req, &resp, true)
	if err != nil {
		return RespGetLeadsInfoByComponentId{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetLeadsInfoByComponentId{}, bizErr
	}
	return resp, nil
}
