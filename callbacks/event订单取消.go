package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单取消
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_cancel.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderCancel{})
}

type ChannelsEcOrderCancel struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		CancelType int64  `json:"cancel_type"`
		OrderID    string `json:"order_id"`
	} `json:"order_info"`
}

func (ChannelsEcOrderCancel) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderCancel) GetEventType() string {
	return "channels_ec_order_cancel"
}

func (m ChannelsEcOrderCancel) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderCancel) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderCancel{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		OrderInfo: struct {
			CancelType int64  `json:"cancel_type"`
			OrderID    string `json:"order_id"`
		}{
			CancelType: gjson.GetBytes(data, "order_info.cancel_type").Int(),
			OrderID:    gjson.GetBytes(data, "order_info.order_id").String(),
		},
	}
	return temp, nil
}
