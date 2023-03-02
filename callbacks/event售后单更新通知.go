package callbacks

import "encoding/json"

// 售后单更新通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/aftersale/ec_callback/channels_ec_aftersale_update.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcAftersaleUpdate{})
}

type ChannelsEcAftersaleUpdate struct {
	CreateTime                      int    `json:"CreateTime"`
	Event                           string `json:"Event"`
	FromUserName                    string `json:"FromUserName"`
	MsgType                         string `json:"MsgType"`
	ToUserName                      string `json:"ToUserName"`
	FinderShopAftersaleStatusUpdate struct {
		AfterSaleOrderID string `json:"after_sale_order_id"`
		OrderID          string `json:"order_id"`
		Status           string `json:"status"`
	} `json:"finder_shop_aftersale_status_update"`
}

func (ChannelsEcAftersaleUpdate) GetMessageType() string {
	return "event"
}

func (ChannelsEcAftersaleUpdate) GetEventType() string {
	return "channels_ec_aftersale_update"
}

func (m ChannelsEcAftersaleUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcAftersaleUpdate) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcAftersaleUpdate
	err := json.Unmarshal(data, &temp)
	return temp, err
}
