package apis

import (
	"encoding/json"
)

// 创建优惠券
// 文档：https://developers.weixin.qq.com/doc/channels/API/coupon/create.html

type ReqCouponCreate struct {
	DiscountInfo CouponDiscountInfo `json:"discount_info"`
	ExtInfo      CouponExtInfo      `json:"ext_info"`
	Name         string             `json:"name"`
	PromoteInfo  CouponPromoteInfo  `json:"promote_info"`
	ReceiveInfo  CouponReceiveInfo  `json:"receive_info"`
	Type         int                `json:"type"`
	ValidInfo    CouponValidInfo    `json:"valid_info"`
}

type CouponDiscountCondition struct {
	ProductCnt   int      `json:"product_cnt"`
	ProductIds   []string `json:"product_ids"`
	ProductPrice int      `json:"product_price"`
}

type CouponDiscountInfo struct {
	DiscountCondition CouponDiscountCondition `json:"discount_condition"`
	DiscountFee       int                     `json:"discount_fee"`
	DiscountNum       int                     `json:"discount_num"`
}

type CouponExtInfo struct {
	InvalidTime   int    `json:"invalid_time"`
	JumpProductID string `json:"jump_product_id"`
	Notes         string `json:"notes"`
	ValidTime     int    `json:"valid_time"`
}

type CouponReceiveInfo struct {
	EndTime           int `json:"end_time"`
	LimitNumOnePerson int `json:"limit_num_one_person"`
	StartTime         int `json:"start_time"`
	TotalNum          int `json:"total_num"`
}

type CouponPromoteInfo struct {
	PromoteType int `json:"promote_type"`
}

type CouponValidInfo struct {
	ValidDayNum int `json:"valid_day_num"`
	ValidType   int `json:"valid_type"`
}

var _ bodyer = ReqCouponCreate{}

func (x ReqCouponCreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCouponCreate struct {
	Data struct {
		CouponID string `json:"coupon_id"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespCouponCreate{}

func (x RespCouponCreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCouponCreate(req ReqCouponCreate) (RespCouponCreate, error) {
	var resp RespCouponCreate
	err := c.executeWXApiPost("/channels/ec/coupon/create", req, &resp, true)
	if err != nil {
		return RespCouponCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCouponCreate{}, bizErr
	}
	return resp, nil
}
