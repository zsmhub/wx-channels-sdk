package apis

import (
	"encoding/json"
)

// 删除品牌资质
// 文档：https://developers.weixin.qq.com/doc/channels/API/brand/delete.html

type ReqBrandDelete struct {
	BrandID string `json:"brand_id"`
}

var _ bodyer = ReqBrandDelete{}

func (x ReqBrandDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespBrandDelete struct {
	CommonResp
}

var _ bodyer = RespBrandDelete{}

func (x RespBrandDelete) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecBrandDelete(req ReqBrandDelete) (RespBrandDelete, error) {
	var resp RespBrandDelete
	err := c.executeWXApiPost("/channels/ec/brand/delete", req, &resp, true)
	if err != nil {
		return RespBrandDelete{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBrandDelete{}, bizErr
	}
	return resp, nil
}
