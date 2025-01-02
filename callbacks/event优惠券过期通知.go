package callbacks

import "encoding/json"

// 优惠券过期通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_expire.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponExpire{})
}

type ChannelsEcCouponExpire struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID   interface{} `json:"coupon_id"`
		ExpireTime interface{} `json:"expire_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponExpire) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponExpire) GetEventType() string {
	return "channels_ec_coupon_expire"
}

func (m ChannelsEcCouponExpire) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponExpire) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcCouponExpire
	err := json.Unmarshal(data, &temp)
	return temp, err
}
