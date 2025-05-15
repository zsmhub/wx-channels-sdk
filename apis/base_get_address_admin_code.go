package apis

import "encoding/json"

// 获取地址编码
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/getaddresscode.html

type ReqAddresscodeGet struct {
	AddrCode int `json:"addr_code"`
}

var _ bodyer = ReqAddresscodeGet{}

func (x ReqAddresscodeGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAddresscodeGet struct {
	CommonResp
	AddrsMsg struct {
		Name  string `json:"name"`
		Code  int    `json:"code"`
		Level int    `json:"level"`
	} `json:"addrs_msg"`
	NextLevelAddrs []struct {
		Name  string `json:"name"`
		Code  int    `json:"code"`
		Level int    `json:"level"`
	} `json:"next_level_addrs"`
}

var _ bodyer = RespAddresscodeGet{}

func (x RespAddresscodeGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAddresscodeGet(req ReqAddresscodeGet) (RespAddresscodeGet, error) {
	var resp RespAddresscodeGet
	err := c.executeWXApiPost("/channels/ec/basics/addresscode/get", req, &resp, true)
	if err != nil {
		return RespAddresscodeGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAddresscodeGet{}, bizErr
	}
	return resp, nil
}
