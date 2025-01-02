package callbacks

import "encoding/json"

// 创建优惠券通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_create.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponCreate{})
}

type ChannelsEcCouponCreate struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID   interface{} `json:"coupon_id"`
		CreateTime interface{} `json:"create_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponCreate) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponCreate) GetEventType() string {
	return "channels_ec_coupon_create"
}

func (m ChannelsEcCouponCreate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponCreate) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcCouponCreate
	err := json.Unmarshal(data, &temp)
	return temp, err
}
