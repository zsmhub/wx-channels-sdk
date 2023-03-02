package callbacks

import "encoding/json"

// 纠纷更新通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/complaint/callback/channels_ec_complaint_update.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcComplaintUpdate{})
}

type ChannelsEcComplaintUpdate struct {
	CreateTime          int    `json:"CreateTime"`
	Event               string `json:"Event"`
	FromUserName        string `json:"FromUserName"`
	MsgType             string `json:"MsgType"`
	ToUserName          string `json:"ToUserName"`
	FinderShopComplaint struct {
		AfterSaleOrderID string `json:"after_sale_order_id"`
		ComplaintID      string `json:"complaint_id"`
		ComplaintStatus  int    `json:"complaint_status"`
		OrderID          string `json:"order_id"`
	} `json:"finder_shop_complaint"`
}

func (ChannelsEcComplaintUpdate) GetMessageType() string {
	return "event"
}

func (ChannelsEcComplaintUpdate) GetEventType() string {
	return "channels_ec_complaint_update"
}

func (m ChannelsEcComplaintUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcComplaintUpdate) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcComplaintUpdate
	err := json.Unmarshal(data, &temp)
	return temp, err
}
