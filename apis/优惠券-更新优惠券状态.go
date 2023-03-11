package apis

import (
	"encoding/json"
)

// 更新优惠券状态
// 文档：https://developers.weixin.qq.com/doc/channels/API/coupon/update_status.html

type ReqCouponUpdateStatus struct {
	CouponID string `json:"coupon_id"`
	Status   int    `json:"status"`
}

var _ bodyer = ReqCouponUpdateStatus{}

func (x ReqCouponUpdateStatus) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCouponUpdateStatus struct {
	CommonResp
}

var _ bodyer = RespCouponUpdateStatus{}

func (x RespCouponUpdateStatus) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCouponUpdateStatus(req ReqCouponUpdateStatus) (RespCouponUpdateStatus, error) {
	var resp RespCouponUpdateStatus
	err := c.executeWXApiPost("/channels/ec/coupon/update_status", req, &resp, true)
	if err != nil {
		return RespCouponUpdateStatus{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCouponUpdateStatus{}, bizErr
	}
	return resp, nil
}
