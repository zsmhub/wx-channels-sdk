package apis

import (
	"encoding/json"
)

// 获取纠纷单
// 文档：https://developers.weixin.qq.com/doc/channels/API/complaint/getcomplaintorder.html

type ReqAftersaleGetcomplaintorder struct {
	ComplaintID string `json:"complaint_id"`
}

var _ bodyer = ReqAftersaleGetcomplaintorder{}

func (x ReqAftersaleGetcomplaintorder) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleGetcomplaintorder struct {
	AfterSaleOrderID string `json:"after_sale_order_id"`
	CommonResp
	History []struct {
		AfterSaleReason int      `json:"after_sale_reason"`
		AfterSaleType   int      `json:"after_sale_type"`
		ComplaintType   int      `json:"complaint_type"`
		Content         string   `json:"content"`
		ItemType        int      `json:"item_type"`
		MediaIDList     []string `json:"media_id_list"`
		PhoneNumber     string   `json:"phone_number"`
		Time            int      `json:"time"`
	} `json:"history"`
	OrderID string `json:"order_id"`
}

var _ bodyer = RespAftersaleGetcomplaintorder{}

func (x RespAftersaleGetcomplaintorder) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleGetcomplaintorder(req ReqAftersaleGetcomplaintorder) (RespAftersaleGetcomplaintorder, error) {
	var resp RespAftersaleGetcomplaintorder
	err := c.executeWXApiPost("/channels/ec/aftersale/getcomplaintorder", req, &resp, true)
	if err != nil {
		return RespAftersaleGetcomplaintorder{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleGetcomplaintorder{}, bizErr
	}
	return resp, nil
}
