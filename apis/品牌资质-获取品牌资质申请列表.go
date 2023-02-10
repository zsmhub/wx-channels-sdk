package apis

import (
	"encoding/json"
)

// 获取品牌资质申请列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/brand/list_get.html

type ReqBrandListGet struct {
	NextKey  string `json:"next_key"`
	PageSize int    `json:"page_size"`
	Status   int    `json:"status"`
}

var _ bodyer = ReqBrandListGet{}

func (x ReqBrandListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespBrandListGet struct {
	Brands []struct {
		ApplicationDetails struct {
			AcceptanceCertification []string `json:"acceptance_certification"`
		} `json:"application_details"`
		BrandID          string `json:"brand_id"`
		ChName           string `json:"ch_name"`
		ClassificationNo string `json:"classification_no"`
		CreateTime       int    `json:"create_time"`
		EnName           string `json:"en_name"`
		GrantDetails     struct {
			BrandOwnerIDPhotos  []interface{} `json:"brand_owner_id_photos"`
			GrantCertifications []interface{} `json:"grant_certifications"`
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
	} `json:"brands"`
	CommonResp
	NextKey  string `json:"next_key"`
	TotalNum int    `json:"total_num"`
}

var _ bodyer = RespBrandListGet{}

func (x RespBrandListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecBrandListGet(req ReqBrandListGet) (RespBrandListGet, error) {
	var resp RespBrandListGet
	err := c.executeWXApiPost("/channels/ec/brand/list/get", req, &resp, true)
	if err != nil {
		return RespBrandListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBrandListGet{}, bizErr
	}
	return resp, nil
}
