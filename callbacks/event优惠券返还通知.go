package callbacks

import "encoding/json"

// 优惠券返还通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_user_coupon_unuse.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcUserCouponUnuse{})
}

type ChannelsEcUserCouponUnuse struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	UseInfo      struct {
		CouponID     string `json:"coupon_id"`
		OrderID      string `json:"order_id"`
		UnuseTime    string `json:"unuse_time"`
		UserCouponID string `json:"user_coupon_id"`
	} `json:"use_info"`
}

func (ChannelsEcUserCouponUnuse) GetMessageType() string {
	return "event"
}

func (ChannelsEcUserCouponUnuse) GetEventType() string {
	return "channels_ec_user_coupon_unuse"
}

func (m ChannelsEcUserCouponUnuse) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcUserCouponUnuse) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcUserCouponUnuse
	err := json.Unmarshal(data, &temp)
	return temp, err
}
