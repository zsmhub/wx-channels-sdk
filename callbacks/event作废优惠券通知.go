package callbacks

import "encoding/json"

// 作废优惠券通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_invalid.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponInvalid{})
}

type ChannelsEcCouponInvalid struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID    interface{} `json:"coupon_id"`
		InvalidTime interface{} `json:"invalid_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponInvalid) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponInvalid) GetEventType() string {
	return "channels_ec_coupon_invalid"
}

func (m ChannelsEcCouponInvalid) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponInvalid) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcCouponInvalid
	err := json.Unmarshal(data, &temp)
	return temp, err
}
