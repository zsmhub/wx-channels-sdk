package apis

import (
	"encoding/json"
)

// 同意售后
// 文档：https://developers.weixin.qq.com/doc/channels/API/aftersale/acceptapply.html

type ReqAftersaleAcceptapply struct {
	AddressID        string `json:"address_id"`
	AfterSaleOrderID string `json:"after_sale_order_id"`
}

var _ bodyer = ReqAftersaleAcceptapply{}

func (x ReqAftersaleAcceptapply) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleAcceptapply struct {
	CommonResp
}

var _ bodyer = RespAftersaleAcceptapply{}

func (x RespAftersaleAcceptapply) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleAcceptapply(req ReqAftersaleAcceptapply) (RespAftersaleAcceptapply, error) {
	var resp RespAftersaleAcceptapply
	err := c.executeWXApiPost("/channels/ec/aftersale/acceptapply", req, &resp, true)
	if err != nil {
		return RespAftersaleAcceptapply{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleAcceptapply{}, bizErr
	}
	return resp, nil
}
