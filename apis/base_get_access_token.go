package apis

import (
	"encoding/json"

	"net/url"
)

// 获取AccessToken
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/getaccesstoken.html

type ReqGetAccessToken struct {
	GrantType string `json:"grant_type"` // 固定填 client_credential，必填
	Appid     string `json:"appid"`      // 小店唯一凭证，即 小店ID，必填
	Secret    string `json:"secret"`     // 小店唯一凭证密钥，必填
}

var _ urlValuer = ReqGetAccessToken{}

func (x ReqGetAccessToken) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

type RespGetAccessToken struct {
	CommonResp
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。
}

var _ bodyer = RespGetAccessToken{}

func (x RespGetAccessToken) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetAccessToken(req ReqGetAccessToken) (RespGetAccessToken, error) {
	var resp RespGetAccessToken
	err := c.executeWXApiGet("/cgi-bin/token", req, &resp, false)
	if err != nil {
		return RespGetAccessToken{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetAccessToken{}, bizErr
	}
	return resp, nil
}
