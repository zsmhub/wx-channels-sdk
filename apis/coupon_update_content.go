package apis

import (
	"encoding/json"
)

// 更新优惠券内容
// 文档：https://developers.weixin.qq.com/doc/channels/API/coupon/update.html

type ReqCouponUpdate struct {
	CouponID     string             `json:"coupon_id"`
	DiscountInfo CouponDiscountInfo `json:"discount_info"`
	ExtInfo      CouponExtInfo      `json:"ext_info"`
	Name         string             `json:"name"`
	PromoteInfo  CouponPromoteInfo  `json:"promote_info"`
	ReceiveInfo  CouponReceiveInfo  `json:"receive_info"`
	Type         int                `json:"type"`
	ValidInfo    CouponValidInfo    `json:"valid_info"`
}

var _ bodyer = ReqCouponUpdate{}

func (x ReqCouponUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCouponUpdate struct {
	Data struct {
		CouponID string `json:"coupon_id"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespCouponUpdate{}

func (x RespCouponUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCouponUpdate(req ReqCouponUpdate) (RespCouponUpdate, error) {
	var resp RespCouponUpdate
	err := c.executeWXApiPost("/channels/ec/coupon/update", req, &resp, true)
	if err != nil {
		return RespCouponUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCouponUpdate{}, bizErr
	}
	return resp, nil
}
