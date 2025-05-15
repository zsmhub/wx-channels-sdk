package apis

import (
	"encoding/json"
)

// 撤回品牌资质审核
// 文档：https://developers.weixin.qq.com/doc/channels/API/brand/audit_cancel.html

type ReqBrandAuditCancel struct {
	AuditID string `json:"audit_id"`
	BrandID string `json:"brand_id"`
}

var _ bodyer = ReqBrandAuditCancel{}

func (x ReqBrandAuditCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespBrandAuditCancel struct {
	CommonResp
}

var _ bodyer = RespBrandAuditCancel{}

func (x RespBrandAuditCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecBrandAuditCancel(req ReqBrandAuditCancel) (RespBrandAuditCancel, error) {
	var resp RespBrandAuditCancel
	err := c.executeWXApiPost("/channels/ec/brand/audit/cancel", req, &resp, true)
	if err != nil {
		return RespBrandAuditCancel{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBrandAuditCancel{}, bizErr
	}
	return resp, nil
}
