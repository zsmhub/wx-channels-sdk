package apis

import (
	"encoding/json"
)

// 获取留资request_id列表详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/leads/get_leads_request_id.html

type ReqGetLeadsInfoByRequestId struct {
	RequestId  string `json:"request_id"`
	LastBuffer string `json:"last_buffer"`
}

var _ bodyer = ReqGetLeadsInfoByRequestId{}

func (x ReqGetLeadsInfoByRequestId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetLeadsInfoByRequestId struct {
	CommonResp
	UserData []struct {
		Phone string `json:"phone"`
	} `json:"user_data"`
	LastBuffer   string `json:"last_buffer"`
	ContinueFlag bool   `json:"continue_flag"`
}

var _ bodyer = RespGetLeadsInfoByRequestId{}

func (x RespGetLeadsInfoByRequestId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetLeadsInfoByRequestId(req ReqGetLeadsInfoByRequestId) (RespGetLeadsInfoByRequestId, error) {
	var resp RespGetLeadsInfoByRequestId
	err := c.executeWXApiPost("/channels/leads/get_leads_info_by_request_id", req, &resp, true)
	if err != nil {
		return RespGetLeadsInfoByRequestId{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetLeadsInfoByRequestId{}, bizErr
	}
	return resp, nil
}
