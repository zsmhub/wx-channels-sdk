package apis

import (
	"encoding/json"
)

// 获取品牌库列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/brand/all_get.html

type ReqBrandAll struct {
	PageSize int `json:"page_size"`
}

var _ bodyer = ReqBrandAll{}

func (x ReqBrandAll) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespBrandAll struct {
	Brands []struct {
		BrandID string `json:"brand_id"`
		ChName  string `json:"ch_name"`
		EnName  string `json:"en_name"`
	} `json:"brands"`
	ContinueFlag bool `json:"continue_flag"`
	CommonResp
	NextKey string `json:"next_key"`
}

var _ bodyer = RespBrandAll{}

func (x RespBrandAll) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecBrandAll(req ReqBrandAll) (RespBrandAll, error) {
	var resp RespBrandAll
	err := c.executeWXApiPost("/channels/ec/brand/all", req, &resp, true)
	if err != nil {
		return RespBrandAll{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBrandAll{}, bizErr
	}
	return resp, nil
}
