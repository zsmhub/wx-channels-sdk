package apis

import (
	"encoding/json"
)

// 获取优惠券详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/coupon/get.html

type ReqCouponGet struct {
	CouponID string `json:"coupon_id"`
}

var _ bodyer = ReqCouponGet{}

func (x ReqCouponGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCouponGet struct {
	CommonResp
	Coupon struct {
		CouponID   string `json:"coupon_id"`
		CouponInfo struct {
			DiscountInfo struct {
				DiscountNum int `json:"discount_num"`
			} `json:"discount_info"`
			Name        string `json:"name"`
			PromoteInfo struct {
				PromoteType int `json:"promote_type"`
			} `json:"promote_info"`
			ReceiveInfo struct {
				EndTime           int `json:"end_time"`
				LimitNumOnePerson int `json:"limit_num_one_person"`
				StartTime         int `json:"start_time"`
				TotalNum          int `json:"total_num"`
			} `json:"receive_info"`
			ValidInfo struct {
				EndTime     int `json:"end_time"`
				StartTime   int `json:"start_time"`
				ValidDayNum int `json:"valid_day_num"`
				ValidType   int `json:"valid_type"`
			} `json:"valid_info"`
		} `json:"coupon_info"`
		CreateTime int `json:"create_time"`
		Status     int `json:"status"`
		StockInfo  struct {
			IssuedNum  int `json:"issued_num"`
			ReceiveNum int `json:"receive_num"`
			UsedNum    int `json:"used_num"`
		} `json:"stock_info"`
		Type       int `json:"type"`
		UpdateTime int `json:"update_time"`
	} `json:"coupon"`
}

var _ bodyer = RespCouponGet{}

func (x RespCouponGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCouponGet(req ReqCouponGet) (RespCouponGet, error) {
	var resp RespCouponGet
	err := c.executeWXApiPost("/channels/ec/coupon/get", req, &resp, true)
	if err != nil {
		return RespCouponGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCouponGet{}, bizErr
	}
	return resp, nil
}
