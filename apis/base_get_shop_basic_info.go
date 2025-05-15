package apis

import (
	"encoding/json"

	"net/url"
)

// 获取店铺基本信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/getbasicinfo.html

type ReqEcBasicsInfoGet struct{}

var _ urlValuer = ReqEcBasicsInfoGet{}

func (x ReqEcBasicsInfoGet) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

type RespEcBasicsInfoGet struct {
	CommonResp
	Info struct {
		Nickname    string `json:"nickname"`     // 店铺名称
		HeadimgURL  string `json:"headimg_url"`  // 店铺头像URL
		SubjectType string `json:"subject_type"` // 店铺类型，目前为"企业"或"个体工商户"
	} `json:"info"`
}

var _ bodyer = RespEcBasicsInfoGet{}

func (x RespEcBasicsInfoGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecEcBasicsInfoGet(req ReqEcBasicsInfoGet) (RespEcBasicsInfoGet, error) {
	var resp RespEcBasicsInfoGet
	err := c.executeWXApiGet("/channels/ec/basics/info/get", req, &resp, true)
	if err != nil {
		return RespEcBasicsInfoGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespEcBasicsInfoGet{}, bizErr
	}
	return resp, nil
}
