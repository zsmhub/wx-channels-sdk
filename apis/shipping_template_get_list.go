package apis

import (
	"encoding/json"
)

// 获取运费模板列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/merchant/getfreighttemplatelist.html

type ReqMerchantGetfreighttemplatelist struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

var _ bodyer = ReqMerchantGetfreighttemplatelist{}

func (x ReqMerchantGetfreighttemplatelist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespMerchantGetfreighttemplatelist struct {
	CommonResp
	TemplateIDList []string `json:"template_id_list"`
}

var _ bodyer = RespMerchantGetfreighttemplatelist{}

func (x RespMerchantGetfreighttemplatelist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMerchantGetfreighttemplatelist(req ReqMerchantGetfreighttemplatelist) (RespMerchantGetfreighttemplatelist, error) {
	var resp RespMerchantGetfreighttemplatelist
	err := c.executeWXApiPost("/channels/ec/merchant/getfreighttemplatelist", req, &resp, true)
	if err != nil {
		return RespMerchantGetfreighttemplatelist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespMerchantGetfreighttemplatelist{}, bizErr
	}
	return resp, nil
}
