package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单确认收货
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_confirm.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderConfirm{})
}

type ChannelsEcOrderConfirm struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		ConfirmType int64  `json:"confirm_type"`
		OrderID     string `json:"order_id"`
	} `json:"order_info"`
}

func (ChannelsEcOrderConfirm) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderConfirm) GetEventType() string {
	return "channels_ec_order_confirm"
}

func (m ChannelsEcOrderConfirm) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderConfirm) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderConfirm{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		OrderInfo: struct {
			ConfirmType int64  `json:"confirm_type"`
			OrderID     string `json:"order_id"`
		}{
			ConfirmType: gjson.GetBytes(data, "order_info.confirm_type").Int(),
			OrderID:     gjson.GetBytes(data, "order_info.order_id").String(),
		},
	}
	return temp, nil
}
