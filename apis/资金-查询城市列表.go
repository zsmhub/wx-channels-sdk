package apis

import (
	"encoding/json"
)

// 查询城市列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/bank/getcity.html

type ReqGetcity struct {
	ProvinceCode int `json:"province_code"`
}

var _ bodyer = ReqGetcity{}

func (x ReqGetcity) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetcity struct {
	CommonResp
	Data []struct {
		BankAddressCode string `json:"bank_address_code"`
		CityCode        int    `json:"city_code"`
		CityName        string `json:"city_name"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

var _ bodyer = RespGetcity{}

func (x RespGetcity) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetcity(req ReqGetcity) (RespGetcity, error) {
	var resp RespGetcity
	err := c.executeWXApiPost("/shop/funds/getcity", req, &resp, true)
	if err != nil {
		return RespGetcity{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetcity{}, bizErr
	}
	return resp, nil
}
