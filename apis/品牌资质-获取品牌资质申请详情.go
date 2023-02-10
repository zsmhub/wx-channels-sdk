package apis

import (
	"encoding/json"
)

// 获取品牌资质申请详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/brand/get.html

type ReqBrandGet struct {
	BrandID string `json:"brand_id"`
}

var _ bodyer = ReqBrandGet{}

func (x ReqBrandGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespBrandGet struct {
	Brand struct {
		ApplicationDetails []interface{} `json:"application_details"`
		BrandID            string        `json:"brand_id"`
		ChName             string        `json:"ch_name"`
		ClassificationNo   string        `json:"classification_no"`
		CreateTime         int           `json:"create_time"`
		EnName             string        `json:"en_name"`
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
		Status          int `json:"status"`
		TradeMarkSymbol int `json:"trade_mark_symbol"`
		UpdateTime      int `json:"update_time"`
	} `json:"brand"`
	CommonResp
}

var _ bodyer = RespBrandGet{}

func (x RespBrandGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecBrandGet(req ReqBrandGet) (RespBrandGet, error) {
	var resp RespBrandGet
	err := c.executeWXApiPost("/channels/ec/brand/get", req, &resp, true)
	if err != nil {
		return RespBrandGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBrandGet{}, bizErr
	}
	return resp, nil
}
