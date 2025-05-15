package apis

import (
	"encoding/json"
)

// 查询开通的网点账号信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/ewaybill/get_account.html

type ReqLogisticsEwaybillBizAccountGet struct {
	Limit       int  `json:"limit"`
	NeedBalance bool `json:"need_balance"`
}

var _ bodyer = ReqLogisticsEwaybillBizAccountGet{}

func (x ReqLogisticsEwaybillBizAccountGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespLogisticsEwaybillBizAccountGet struct {
	AccountList []struct {
		AcctID     string `json:"acct_id"`
		Allocated  int    `json:"allocated"`
		Available  int    `json:"available"`
		Cancel     int    `json:"cancel"`
		DeliveryID string `json:"delivery_id"`
		Recycled   int    `json:"recycled"`
		ShopID     string `json:"shop_id"`
		SiteInfo   struct {
			Address struct {
				CityCode      string `json:"city_code"`
				CityName      string `json:"city_name"`
				DetailAddress string `json:"detail_address"`
				DistrictCode  string `json:"district_code"`
				DistrictName  string `json:"district_name"`
				ProvinceCode  string `json:"province_code"`
				ProvinceName  string `json:"province_name"`
				StreetCode    string `json:"street_code"`
				StreetName    string `json:"street_name"`
			} `json:"address"`
			Contact struct {
				Mobile string `json:"mobile"`
				Name   string `json:"name"`
				Phone  string `json:"phone"`
			} `json:"contact"`
			DeliveryID   string `json:"delivery_id"`
			SiteCode     string `json:"site_code"`
			SiteFullname string `json:"site_fullname"`
			SiteName     string `json:"site_name"`
			SiteStatus   int    `json:"site_status"`
		} `json:"site_info"`
		Status int `json:"status"`
	} `json:"account_list"`
	CommonResp
	TotalNum int `json:"total_num"`
}

var _ bodyer = RespLogisticsEwaybillBizAccountGet{}

func (x RespLogisticsEwaybillBizAccountGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLogisticsEwaybillBizAccountGet(req ReqLogisticsEwaybillBizAccountGet) (RespLogisticsEwaybillBizAccountGet, error) {
	var resp RespLogisticsEwaybillBizAccountGet
	err := c.executeWXApiPost("/channels/ec/logistics/ewaybill/biz/account/get", req, &resp, true)
	if err != nil {
		return RespLogisticsEwaybillBizAccountGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLogisticsEwaybillBizAccountGet{}, bizErr
	}
	return resp, nil
}
