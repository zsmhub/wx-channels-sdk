package apis

import (
	"encoding/json"
	"net/url"
)

// 获取账号申请通过的类目和资质信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/category/getavailablebizcat.html

type ReqCategoryListGet struct{}

var _ urlValuer = ReqCategoryListGet{}

func (x ReqCategoryListGet) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

type RespCategoryListGet struct {
	CommonResp
	List []struct {
		CatId string `json:"cat_id"`
		QuaId string `json:"qua_id"`
	} `json:"list"`
}

var _ bodyer = RespCategoryListGet{}

func (x RespCategoryListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCategoryListGet(req ReqCategoryListGet) (RespCategoryListGet, error) {
	var resp RespCategoryListGet
	err := c.executeWXApiGet("/channels/ec/category/list/get", req, &resp, true)
	if err != nil {
		return RespCategoryListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCategoryListGet{}, bizErr
	}
	return resp, nil
}
