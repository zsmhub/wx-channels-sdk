package apis

import (
	"encoding/json"
)

// 获取用户优惠券详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/coupon/get_user_coupon.html

type ReqCouponGetUserCoupon struct {
	Openid       string `json:"openid"`
	UserCouponID string `json:"user_coupon_id"`
}

var _ bodyer = ReqCouponGetUserCoupon{}

func (x ReqCouponGetUserCoupon) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCouponGetUserCoupon struct {
	CommonResp
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	UserCoupon struct {
		UserCouponID string `json:"user_coupon_id"`
		CouponID     string `json:"coupon_id"`
		Status       int    `json:"status"`
		CreateTime   int    `json:"create_time"`
		UpdateTime   int    `json:"update_time"`
		StartTime    int    `json:"start_time"`
		EndTime      int    `json:"end_time"`
		ExtInfo      struct {
			UseTime int `json:"use_time"`
		} `json:"ext_info"`
		OrderID     string `json:"order_id"`
		DiscountFee int    `json:"discount_fee"`
	} `json:"user_coupon"`
}

var _ bodyer = RespCouponGetUserCoupon{}

func (x RespCouponGetUserCoupon) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCouponGetUserCoupon(req ReqCouponGetUserCoupon) (RespCouponGetUserCoupon, error) {
	var resp RespCouponGetUserCoupon
	err := c.executeWXApiPost("/channels/ec/coupon/get_user_coupon", req, &resp, true)
	if err != nil {
		return RespCouponGetUserCoupon{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCouponGetUserCoupon{}, bizErr
	}
	return resp, nil
}
