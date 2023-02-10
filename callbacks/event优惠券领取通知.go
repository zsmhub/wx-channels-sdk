package callbacks

import "encoding/json"

// 优惠券领取通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_receive.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponReceive{})
}

type ChannelsEcCouponReceive struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	ReceiveInfo  struct {
		CouponID     string `json:"coupon_id"`
		ReceiveTime  string `json:"receive_time"`
		UserCouponID string `json:"user_coupon_id"`
	} `json:"receive_info"`
}

func (ChannelsEcCouponReceive) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponReceive) GetEventType() string {
	return "channels_ec_coupon_receive"
}

func (m ChannelsEcCouponReceive) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponReceive) ParseFromJson(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ChannelsEcCouponReceive
	err := json.Unmarshal(data, &temp)
	return temp, err
}
