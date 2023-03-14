package callbacks

import "github.com/tidwall/gjson"

// 订单结算成功
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_settle.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderSettle{})
}

type ChannelsEcOrderSettle struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID    string `json:"order_id"`
		SettleTime int64  `json:"settle_time"`
	} `json:"order_info"`
}

func (ChannelsEcOrderSettle) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderSettle) GetEventType() string {
	return "channels_ec_order_settle"
}

func (m ChannelsEcOrderSettle) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderSettle) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderSettle{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		OrderInfo: struct {
			OrderID    string `json:"order_id"`
			SettleTime int64  `json:"settle_time"`
		}{
			OrderID:    gjson.GetBytes(data, "order_info.order_id").String(),
			SettleTime: gjson.GetBytes(data, "order_info.settle_time").Int(),
		},
	}
	return temp, nil
}
