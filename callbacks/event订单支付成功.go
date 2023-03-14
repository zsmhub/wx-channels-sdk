package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单支付成功
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_pay.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderPay{})
}

type ChannelsEcOrderPay struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID string `json:"order_id"`
		PayTime int64  `json:"pay_time"`
	} `json:"order_info"`
}

func (ChannelsEcOrderPay) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderPay) GetEventType() string {
	return "channels_ec_order_pay"
}

func (m ChannelsEcOrderPay) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderPay) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderPay{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		OrderInfo: struct {
			OrderID string `json:"order_id"`
			PayTime int64  `json:"pay_time"`
		}{
			OrderID: gjson.GetBytes(data, "order_info.order_id").String(),
			PayTime: gjson.GetBytes(data, "order_info.pay_time").Int(),
		},
	}
	return temp, nil
}
