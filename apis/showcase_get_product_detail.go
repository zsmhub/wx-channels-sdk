package apis

import (
	"encoding/json"
)

// 获取橱窗商品详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/windowproduct/get.html

type ReqWindowProductGet struct {
	Appid     string `json:"appid"`
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqWindowProductGet{}

func (x ReqWindowProductGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespWindowProductGet struct {
	CommonResp
	Product struct {
		Appid             string `json:"appid"`
		Banned            bool   `json:"banned"`
		ImgURL            string `json:"img_url"`
		IsHideForWindow   int    `json:"is_hide_for_window"`
		LimitDiscountInfo struct {
			DiscountPrice int  `json:"discount_price"`
			EndTimeMs     int  `json:"end_time_ms"`
			IsEffect      bool `json:"is_effect"`
			Stock         int  `json:"stock"`
		} `json:"limit_discount_info"`
		MarketPrice  int    `json:"market_price"`
		OutProductID string `json:"out_product_id"`
		PagePath     struct {
			Appid        string `json:"appid"`
			FullPagePath string `json:"full_page_path"`
			HalfPagePath string `json:"half_page_path"`
		} `json:"page_path"`
		PlatformID              int    `json:"platform_id"`
		PlatformName            string `json:"platform_name"`
		ProductID               string `json:"product_id"`
		ProductRecommendWording string `json:"product_recommend_wording"`
		Sales                   int    `json:"sales"`
		SellingPrice            int    `json:"selling_price"`
		Status                  int    `json:"status"`
		Stock                   int    `json:"stock"`
		ThirdCategoryID         int    `json:"third_category_id"`
		Title                   string `json:"title"`
	} `json:"product"`
}

var _ bodyer = RespWindowProductGet{}

func (x RespWindowProductGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecWindowProductGet(req ReqWindowProductGet) (RespWindowProductGet, error) {
	var resp RespWindowProductGet
	err := c.executeWXApiPost("/channels/ec/window/product/get", req, &resp, true)
	if err != nil {
		return RespWindowProductGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespWindowProductGet{}, bizErr
	}
	return resp, nil
}
