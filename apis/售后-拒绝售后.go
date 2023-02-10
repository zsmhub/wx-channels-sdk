package apis

import (
	"encoding/json"
)

// 拒绝售后
// 文档：https://developers.weixin.qq.com/doc/channels/API/aftersale/rejectapply.html

type ReqAftersaleRejectapply struct {
	AfterSaleOrderID string `json:"after_sale_order_id"`
	RejectReason     string `json:"reject_reason"`
}

var _ bodyer = ReqAftersaleRejectapply{}

func (x ReqAftersaleRejectapply) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleRejectapply struct {
	CommonResp
}

var _ bodyer = RespAftersaleRejectapply{}

func (x RespAftersaleRejectapply) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleRejectapply(req ReqAftersaleRejectapply) (RespAftersaleRejectapply, error) {
	var resp RespAftersaleRejectapply
	err := c.executeWXApiPost("/channels/ec/aftersale/rejectapply", req, &resp, true)
	if err != nil {
		return RespAftersaleRejectapply{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleRejectapply{}, bizErr
	}
	return resp, nil
}
