package apis

import (
	"encoding/json"
)

// 取消类目提审
// 文档：https://developers.weixin.qq.com/doc/channels/API/category/cancelauditcategory.html

type ReqCategoryAuditCancel struct {
	AuditID string `json:"audit_id"`
}

var _ bodyer = ReqCategoryAuditCancel{}

func (x ReqCategoryAuditCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCategoryAuditCancel struct {
	CommonResp
}

var _ bodyer = RespCategoryAuditCancel{}

func (x RespCategoryAuditCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCategoryAuditCancel(req ReqCategoryAuditCancel) (RespCategoryAuditCancel, error) {
	var resp RespCategoryAuditCancel
	err := c.executeWXApiPost("/channels/ec/category/audit/cancel", req, &resp, true)
	if err != nil {
		return RespCategoryAuditCancel{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCategoryAuditCancel{}, bizErr
	}
	return resp, nil
}
