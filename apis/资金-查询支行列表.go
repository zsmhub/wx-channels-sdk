package apis

import (
	"encoding/json"
)

// 查询支行列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/funds/bank/getsubbranch.html

type ReqGetsubbranch struct {
	BankCode string `json:"bank_code"`
	CityCode string `json:"city_code"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

var _ bodyer = ReqGetsubbranch{}

func (x ReqGetsubbranch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetsubbranch struct {
	CommonResp
	AccountBank     string `json:"account_bank"`
	AccountBankCode int    `json:"account_bank_code"`
	BankAlias       string `json:"bank_alias"`
	BankAliasCode   string `json:"bank_alias_code"`
	Count           int    `json:"count"`
	Data            []struct {
		BranchID   string `json:"branch_id"`
		BranchName string `json:"branch_name"`
	} `json:"data"`
	TotalCount int `json:"total_count"`
}

var _ bodyer = RespGetsubbranch{}

func (x RespGetsubbranch) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetsubbranch(req ReqGetsubbranch) (RespGetsubbranch, error) {
	var resp RespGetsubbranch
	err := c.executeWXApiPost("/shop/funds/getsubbranch", req, &resp, true)
	if err != nil {
		return RespGetsubbranch{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetsubbranch{}, bizErr
	}
	return resp, nil
}
