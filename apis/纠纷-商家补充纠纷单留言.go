package apis

import (
	"encoding/json"
)

// 商家补充纠纷单留言
// 文档：https://developers.weixin.qq.com/doc/channels/API/complaint/addcomplaintmaterial.html

type ReqAftersaleAddcomplaintmaterial struct {
	ComplaintID string   `json:"complaint_id"`
	Content     string   `json:"content"`
	MediaIDList []string `json:"media_id_list"`
}

var _ bodyer = ReqAftersaleAddcomplaintmaterial{}

func (x ReqAftersaleAddcomplaintmaterial) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleAddcomplaintmaterial struct {
	CommonResp
}

var _ bodyer = RespAftersaleAddcomplaintmaterial{}

func (x RespAftersaleAddcomplaintmaterial) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleAddcomplaintmaterial(req ReqAftersaleAddcomplaintmaterial) (RespAftersaleAddcomplaintmaterial, error) {
	var resp RespAftersaleAddcomplaintmaterial
	err := c.executeWXApiPost("/channels/ec/aftersale/addcomplaintmaterial", req, &resp, true)
	if err != nil {
		return RespAftersaleAddcomplaintmaterial{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleAddcomplaintmaterial{}, bizErr
	}
	return resp, nil
}
