package apis

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

// 通过 mediaid 获取数据
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/getdatabymediaid.html

type ReqMediaGet struct {
	MediaId string `json:"media_id"`
}

var _ urlValuer = ReqMediaGet{}

func (x ReqMediaGet) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

type RespMediaGet struct {
	CommonResp
}

var _ bodyer = RespMediaGet{}

func (x RespMediaGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecMediaGet(req ReqMediaGet) ([]byte, error) {
	byteResp, err := c.executeWXApiGetReturnByte("/channels/ec/basics/media/get", req, true)
	if err != nil {
		return byteResp, err
	}

	if strings.Contains(string(byteResp), "errcode") {
		var r RespMediaGet
		err = json.Unmarshal(byteResp, &r)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(r.ErrMsg)
	}

	return byteResp, nil
}
