package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单其他信息更新
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_ext_info.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderExtInfoUpdate{})
}

type ChannelsEcOrderExtInfoUpdate struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID string `json:"order_id"`
		Type    int64  `json:"type"`
	} `json:"order_info"`
}

func (ChannelsEcOrderExtInfoUpdate) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderExtInfoUpdate) GetEventType() string {
	return "channels_ec_order_ext_info_update"
}

func (m ChannelsEcOrderExtInfoUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderExtInfoUpdate) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcOrderExtInfoUpdate{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		OrderInfo: struct {
			OrderID string `json:"order_id"`
			Type    int64  `json:"type"`
		}{
			OrderID: gjson.GetBytes(data, "order_info.order_id").String(),
			Type:    gjson.GetBytes(data, "order_info.type").Int(),
		},
	}
	return temp, nil
}
