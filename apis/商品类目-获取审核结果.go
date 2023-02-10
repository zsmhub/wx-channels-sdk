package apis

import (
	"encoding/json"
)

// 获取审核结果
// 文档：https://developers.weixin.qq.com/doc/channels/API/category/audit_get.html

type ReqCategoryAuditGet struct {
	AuditID string `json:"audit_id"`
}

var _ bodyer = ReqCategoryAuditGet{}

func (x ReqCategoryAuditGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCategoryAuditGet struct {
	Data struct {
		RejectReason string `json:"reject_reason"`
		Status       int    `json:"status"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespCategoryAuditGet{}

func (x RespCategoryAuditGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCategoryAuditGet(req ReqCategoryAuditGet) (RespCategoryAuditGet, error) {
	var resp RespCategoryAuditGet
	err := c.executeWXApiPost("/channels/ec/category/audit/get", req, &resp, true)
	if err != nil {
		return RespCategoryAuditGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCategoryAuditGet{}, bizErr
	}
	return resp, nil
}
