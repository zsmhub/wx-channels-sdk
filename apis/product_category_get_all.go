package apis

import (
	"encoding/json"
	"net/url"
)

// 获取所有的类目
// 文档：https://developers.weixin.qq.com/doc/channels/API/category/getallcategory.html

type ReqCategoryAll struct{}

var _ urlValuer = ReqCategoryAll{}

func (x ReqCategoryAll) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

type RespCategoryAll struct {
	Cats []struct {
		CatAndQua []struct {
			Cat struct {
				CatID  string `json:"cat_id"`
				FCatID string `json:"f_cat_id"`
				Level  int    `json:"level"`
				Name   string `json:"name"`
			} `json:"cat"`
			ProductQua struct {
				Mandatory   bool   `json:"mandatory"`
				NeedToApply bool   `json:"need_to_apply"`
				QuaID       string `json:"qua_id"`
				Tips        string `json:"tips"`
			} `json:"product_qua"`
			Qua struct {
				Mandatory   bool   `json:"mandatory"`
				NeedToApply bool   `json:"need_to_apply"`
				QuaID       string `json:"qua_id"`
				Tips        string `json:"tips"`
			} `json:"qua"`
		} `json:"cat_and_qua"`
	} `json:"cats"`
	CommonResp
}

var _ bodyer = RespCategoryAll{}

func (x RespCategoryAll) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCategoryAll(req ReqCategoryAll) (RespCategoryAll, error) {
	var resp RespCategoryAll
	err := c.executeWXApiGet("/channels/ec/category/all", req, &resp, true)
	if err != nil {
		return RespCategoryAll{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCategoryAll{}, bizErr
	}
	return resp, nil
}
