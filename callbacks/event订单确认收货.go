package callbacks

import "encoding/json"

// 订单确认收货
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_confirm.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderConfirm{})
}

type ChannelsEcOrderConfirm struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		ConfirmType int `json:"confirm_type"`
		OrderID     int `json:"order_id"`
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
	var temp ChannelsEcOrderConfirm
	err := json.Unmarshal(data, &temp)
	return temp, err
}
