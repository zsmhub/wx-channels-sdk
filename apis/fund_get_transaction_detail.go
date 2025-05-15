package apis

import (
	"encoding/json"
)

// 获取资金流水详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/getfundsflowdetail.html

type ReqFundsGetfundsflowdetail struct {
	FlowID string `json:"flow_id"`
}

var _ bodyer = ReqFundsGetfundsflowdetail{}

func (x ReqFundsGetfundsflowdetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFundsGetfundsflowdetail struct {
	CommonResp
	FundsFlow struct {
		Amount          int    `json:"amount"`
		Balance         int    `json:"balance"`
		FlowID          string `json:"flow_id"`
		FlowType        int    `json:"flow_type"`
		FundsType       int    `json:"funds_type"`
		RelatedInfoList []struct {
			AftersaleID   string `json:"aftersale_id"`
			OrderID       string `json:"order_id"`
			RelatedType   int    `json:"related_type"`
			TransactionID string `json:"transaction_id"`
		} `json:"related_info_list"`
	} `json:"funds_flow"`
}

var _ bodyer = RespFundsGetfundsflowdetail{}

func (x RespFundsGetfundsflowdetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFundsGetfundsflowdetail(req ReqFundsGetfundsflowdetail) (RespFundsGetfundsflowdetail, error) {
	var resp RespFundsGetfundsflowdetail
	err := c.executeWXApiPost("/channels/ec/funds/getfundsflowdetail", req, &resp, true)
	if err != nil {
		return RespFundsGetfundsflowdetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFundsGetfundsflowdetail{}, bizErr
	}
	return resp, nil
}
