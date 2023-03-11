package apis

import (
	"encoding/json"
)

// 获取用户优惠券ID列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/coupon/get_user_coupon_list.html

type ReqCouponGetUserCouponList struct {
	Openid   string `json:"openid"`
	Page     int    `json:"page"`
	PageCtx  string `json:"page_ctx"`
	PageSize int    `json:"page_size"`
	Status   int    `json:"status"`
}

var _ bodyer = ReqCouponGetUserCouponList{}

func (x ReqCouponGetUserCouponList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCouponGetUserCouponList struct {
	CommonResp
	PageCtx        string `json:"page_ctx"`
	TotalNum       int    `json:"total_num"`
	UserCouponList []struct {
		CouponID     string `json:"coupon_id"`
		UserCouponID string `json:"user_coupon_id"`
	} `json:"user_coupon_list"`
}

var _ bodyer = RespCouponGetUserCouponList{}

func (x RespCouponGetUserCouponList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCouponGetUserCouponList(req ReqCouponGetUserCouponList) (RespCouponGetUserCouponList, error) {
	var resp RespCouponGetUserCouponList
	err := c.executeWXApiPost("/channels/ec/coupon/get_user_coupon_list", req, &resp, true)
	if err != nil {
		return RespCouponGetUserCouponList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCouponGetUserCouponList{}, bizErr
	}
	return resp, nil
}
