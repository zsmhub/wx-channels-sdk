package apis

import (
	"encoding/json"
)

// 撤回商品审核
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/audit_cancel.html

type ReqProductAuditCancel struct {
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqProductAuditCancel{}

func (x ReqProductAuditCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductAuditCancel struct {
	CommonResp
}

var _ bodyer = RespProductAuditCancel{}

func (x RespProductAuditCancel) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductAuditCancel(req ReqProductAuditCancel) (RespProductAuditCancel, error) {
	var resp RespProductAuditCancel
	err := c.executeWXApiPost("/channels/ec/product/audit/cancel", req, &resp, true)
	if err != nil {
		return RespProductAuditCancel{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductAuditCancel{}, bizErr
	}
	return resp, nil
}
