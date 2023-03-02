package callbacks

import "encoding/json"

// 订单支付成功
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_pay.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderPay{})
}

type ChannelsEcOrderPay struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID int `json:"order_id"`
		PayTime int `json:"pay_time"`
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
	var temp ChannelsEcOrderPay
	err := json.Unmarshal(data, &temp)
	return temp, err
}
