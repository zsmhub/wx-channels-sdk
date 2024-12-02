package callbacks

import "encoding/json"

// 优惠券核销通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_user_coupon_use.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcUserCouponUse{})
}

type ChannelsEcUserCouponUse struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	UseInfo      struct {
		CouponID     string `json:"coupon_id"`
		OrderID      string `json:"order_id"`
		UseTime      string `json:"use_time"`
		UserCouponID string `json:"user_coupon_id"`
	} `json:"use_info"`
}

func (ChannelsEcUserCouponUse) GetMessageType() string {
	return "event"
}

func (ChannelsEcUserCouponUse) GetEventType() string {
	return "channels_ec_user_coupon_use"
}

func (m ChannelsEcUserCouponUse) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcUserCouponUse) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcUserCouponUse
	err := json.Unmarshal(data, &temp)
	return temp, err
}
