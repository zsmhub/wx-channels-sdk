package apis

import (
	"encoding/json"
)

// 查询大陆银行省份列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/bank/getprovince.html

type ReqGetprovince struct{}

var _ bodyer = ReqGetprovince{}

func (x ReqGetprovince) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetprovince struct {
	CommonResp
	Data []struct {
		ProvinceCode int    `json:"province_code"`
		ProvinceName string `json:"province_name"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

var _ bodyer = RespGetprovince{}

func (x RespGetprovince) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetprovince(req ReqGetprovince) (RespGetprovince, error) {
	var resp RespGetprovince
	err := c.executeWXApiPost("/shop/funds/getprovince", req, &resp, true)
	if err != nil {
		return RespGetprovince{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetprovince{}, bizErr
	}
	return resp, nil
}
