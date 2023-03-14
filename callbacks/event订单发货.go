package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单发货
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_deliver.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderDeliver{})
}

type ChannelsEcOrderDeliver struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		FinishDelivery int64  `json:"finish_delivery"`
		OrderID        string `json:"order_id"`
	} `json:"order_info"`
}

func (ChannelsEcOrderDeliver) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderDeliver) GetEventType() string {
	return "channels_ec_order_deliver"
}

func (m ChannelsEcOrderDeliver) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderDeliver) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderDeliver{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		OrderInfo: struct {
			FinishDelivery int64  `json:"finish_delivery"`
			OrderID        string `json:"order_id"`
		}{
			FinishDelivery: gjson.GetBytes(data, "order_info.finish_delivery").Int(),
			OrderID:        gjson.GetBytes(data, "order_info.order_id").String(),
		},
	}
	return temp, nil
}
