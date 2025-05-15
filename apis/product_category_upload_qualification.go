package apis

import (
	"encoding/json"
)

// 上传类目资质
// 文档：https://developers.weixin.qq.com/doc/channels/API/category/add.html

type ReqCategoryAdd struct {
	CategoryInfo struct {
		Certificate []string `json:"certificate"`
		Level1      int      `json:"level1"`
		Level2      int      `json:"level2"`
		Level3      int      `json:"level3"`
	} `json:"category_info"`
}

var _ bodyer = ReqCategoryAdd{}

func (x ReqCategoryAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCategoryAdd struct {
	AuditID string `json:"audit_id"`
	CommonResp
}

var _ bodyer = RespCategoryAdd{}

func (x RespCategoryAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCategoryAdd(req ReqCategoryAdd) (RespCategoryAdd, error) {
	var resp RespCategoryAdd
	err := c.executeWXApiPost("/channels/ec/category/add", req, &resp, true)
	if err != nil {
		return RespCategoryAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCategoryAdd{}, bizErr
	}
	return resp, nil
}
