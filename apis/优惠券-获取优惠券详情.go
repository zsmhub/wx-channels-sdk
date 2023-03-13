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
		Type       int    `json:"type"`
		Status     int    `json:"status"`
		CreateTime int    `json:"create_time"`
		UpdateTime int    `json:"update_time"`
		CouponInfo struct {
			Name      string `json:"name"`
			ValidInfo struct {
				ValidType   int `json:"valid_type"`
				ValidDayNum int `json:"valid_day_num"`
				StartTime   int `json:"start_time"`
				EndTime     int `json:"end_time"`
			} `json:"valid_info"`
			PromoteInfo struct {
				PromoteType int `json:"promote_type"`
			} `json:"promote_info"`
			DiscountInfo struct {
				DiscountNum       int `json:"discount_num"`
				DiscountFee       int `json:"discount_fee"`
				DiscountCondition struct {
					ProductCnt   int      `json:"product_cnt"`
					ProductPrice int      `json:"product_price"`
					ProductIds   []string `json:"product_ids"`
				} `json:"discount_condition"`
			} `json:"discount_info"`
			ExtInfo struct {
				InvalidTime   int    `json:"invalid_time"`
				JumpProductId int    `json:"jump_product_id"`
				Notes         string `json:"notes"`
				ValidTime     int    `json:"valid_time"`
			}
			ReceiveInfo struct {
				EndTime           int `json:"end_time"`
				LimitNumOnePerson int `json:"limit_num_one_person"`
				StartTime         int `json:"start_time"`
				TotalNum          int `json:"total_num"`
			} `json:"receive_info"`
		} `json:"coupon_info"`
		StockInfo struct {
			IssuedNum  int `json:"issued_num"`
			ReceiveNum int `json:"receive_num"`
			UsedNum    int `json:"used_num"`
		} `json:"stock_info"`
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
