package callbacks

import "encoding/json"

// 优惠券信息更新通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_info_change.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponInfoChange{})
}

type ChannelsEcCouponInfoChange struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		ChangeTime interface{} `json:"change_time"`
		CouponID   interface{} `json:"coupon_id"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponInfoChange) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponInfoChange) GetEventType() string {
	return "channels_ec_coupon_info_change"
}

func (m ChannelsEcCouponInfoChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponInfoChange) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcCouponInfoChange
	err := json.Unmarshal(data, &temp)
	return temp, err
}
