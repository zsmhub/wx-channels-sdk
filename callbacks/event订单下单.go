package callbacks

import "encoding/json"

// 订单下单
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_new.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderNew{})
}

type ChannelsEcOrderNew struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID int `json:"order_id"`
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

func (ChannelsEcOrderNew) ParseFromJson(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ChannelsEcOrderNew
	err := json.Unmarshal(data, &temp)
	return temp, err
}
