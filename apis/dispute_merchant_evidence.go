package apis

import (
	"encoding/json"
)

// 商家举证
// 文档：https://developers.weixin.qq.com/doc/channels/API/complaint/addcomplaintproof.html

type ReqAftersaleAddcomplaintproof struct {
	ComplaintID string   `json:"complaint_id"`
	Content     string   `json:"content"`
	MediaIDList []string `json:"media_id_list"`
}

var _ bodyer = ReqAftersaleAddcomplaintproof{}

func (x ReqAftersaleAddcomplaintproof) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleAddcomplaintproof struct {
	CommonResp
}

var _ bodyer = RespAftersaleAddcomplaintproof{}

func (x RespAftersaleAddcomplaintproof) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleAddcomplaintproof(req ReqAftersaleAddcomplaintproof) (RespAftersaleAddcomplaintproof, error) {
	var resp RespAftersaleAddcomplaintproof
	err := c.executeWXApiPost("/channels/ec/aftersale/addcomplaintproof", req, &resp, true)
	if err != nil {
		return RespAftersaleAddcomplaintproof{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleAddcomplaintproof{}, bizErr
	}
	return resp, nil
}
