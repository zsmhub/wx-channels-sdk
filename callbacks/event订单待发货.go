package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单待发货
// 文档: https://developers.weixin.qq.com/doc/store/shop/notify/order_callback/channels_ec_order_wait_shipping.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderWait{})
}

type ChannelsEcOrderWait struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID string `json:"order_id"`
	} `json:"order_info"`
}

func (ChannelsEcOrderWait) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderWait) GetEventType() string {
	return "channels_ec_order_wait_shipping"
}

func (m ChannelsEcOrderWait) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderWait) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderWait{
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
