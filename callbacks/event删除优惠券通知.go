package callbacks

import "encoding/json"

// 删除优惠券通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_delete.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponDelete{})
}

type ChannelsEcCouponDelete struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID   interface{} `json:"coupon_id"`
		DeleteTime interface{} `json:"delete_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponDelete) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponDelete) GetEventType() string {
	return "channels_ec_coupon_delete"
}

func (m ChannelsEcCouponDelete) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponDelete) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcCouponDelete
	err := json.Unmarshal(data, &temp)
	return temp, err
}
