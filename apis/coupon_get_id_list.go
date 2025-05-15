package apis

import (
	"encoding/json"
)

// 获取优惠券ID列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/coupon/get_list.html

type ReqCouponGetList struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Status   int `json:"status"`
}

var _ bodyer = ReqCouponGetList{}

func (x ReqCouponGetList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCouponGetList struct {
	Coupons []struct {
		CouponID string `json:"coupon_id"`
	} `json:"coupons"`
	CommonResp
	PageCtx  string `json:"page_ctx"`
	TotalNum int    `json:"total_num"`
}

var _ bodyer = RespCouponGetList{}

func (x RespCouponGetList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCouponGetList(req ReqCouponGetList) (RespCouponGetList, error) {
	var resp RespCouponGetList
	err := c.executeWXApiPost("/channels/ec/coupon/get_list", req, &resp, true)
	if err != nil {
		return RespCouponGetList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCouponGetList{}, bizErr
	}
	return resp, nil
}
