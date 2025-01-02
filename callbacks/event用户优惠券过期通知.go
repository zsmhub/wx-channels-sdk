package callbacks

import "encoding/json"

// 用户优惠券过期通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_user_coupon_expire.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcUserCouponExpire{})
}

type ChannelsEcUserCouponExpire struct {
	CreateTime     int    `json:"CreateTime"`
	Event          string `json:"Event"`
	FromUserName   string `json:"FromUserName"`
	MsgType        string `json:"MsgType"`
	ToUserName     string `json:"ToUserName"`
	UserCouponInfo struct {
		CouponID     interface{} `json:"coupon_id"`
		ExpireTime   interface{} `json:"expire_time"`
		UserCouponID interface{} `json:"user_coupon_id"`
	} `json:"user_coupon_info"`
}

func (ChannelsEcUserCouponExpire) GetMessageType() string {
	return "event"
}

func (ChannelsEcUserCouponExpire) GetEventType() string {
	return "channels_ec_user_coupon_expire"
}

func (m ChannelsEcUserCouponExpire) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcUserCouponExpire) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcUserCouponExpire
	err := json.Unmarshal(data, &temp)
	return temp, err
}
