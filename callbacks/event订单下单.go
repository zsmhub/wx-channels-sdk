package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单下单
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_new.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderNew{})
}

type ChannelsEcOrderNew struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID string `json:"order_id"`
	} `json:"order_info"`
}

func (ChannelsEcOrderNew) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderNew) GetEventType() string {
	return "channels_ec_order_new"
}

func (m ChannelsEcOrderNew) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderNew) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderNew{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		OrderInfo: struct {
			OrderID string `json:"order_id"`
		}{
			OrderID: gjson.GetBytes(data, "order_info.order_id").String(),
		},
	}
	return temp, nil
}
