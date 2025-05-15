package apis

import (
	"encoding/json"
)

// 新增品牌资质
// 文档：https://developers.weixin.qq.com/doc/channels/API/brand/add.html

type ReqBrandAdd struct {
	Brand struct {
		ApplicationDetails struct{} `json:"application_details"`
		BrandID            string   `json:"brand_id"`
		ChName             string   `json:"ch_name"`
		ClassificationNo   string   `json:"classification_no"`
		EnName             string   `json:"en_name"`
		GrantDetails       struct {
			BrandOwnerIDPhotos  []string `json:"brand_owner_id_photos"`
			EndTime             int      `json:"end_time"`
			GrantCertifications []string `json:"grant_certifications"`
			GrantLevel          int      `json:"grant_level"`
			IsPermanent         bool     `json:"is_permanent"`
			StartTime           int      `json:"start_time"`
		} `json:"grant_details"`
		GrantType       int `json:"grant_type"`
		RegisterDetails struct {
			EndTime                int      `json:"end_time"`
			IsPermanent            bool     `json:"is_permanent"`
			RegisterCertifications []string `json:"register_certifications"`
			RegisterNo             string   `json:"register_no"`
			Registrant             string   `json:"registrant"`
			RenewCertifications    []string `json:"renew_certifications"`
			StartTime              int      `json:"start_time"`
		} `json:"register_details"`
		TradeMarkSymbol int `json:"trade_mark_symbol"`
	} `json:"brand"`
}

var _ bodyer = ReqBrandAdd{}

func (x ReqBrandAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespBrandAdd struct {
	AuditID string `json:"audit_id"`
	CommonResp
}

var _ bodyer = RespBrandAdd{}

func (x RespBrandAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecBrandAdd(req ReqBrandAdd) (RespBrandAdd, error) {
	var resp RespBrandAdd
	err := c.executeWXApiPost("/channels/ec/brand/add", req, &resp, true)
	if err != nil {
		return RespBrandAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBrandAdd{}, bizErr
	}
	return resp, nil
}
