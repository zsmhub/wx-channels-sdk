package callbacks

import (
	"encoding/json"
)

// 商品上下架
// 文档: https://developers.weixin.qq.com/doc/channels/API/product/callback/ProductSpuListing.html

func init() {
	supportCallback(EventProductSpuListing{})
}

type EventProductSpuListing struct {
	ToUserName      string `json:"ToUserName"`
	FromUserName    string `json:"FromUserName"`
	CreateTime      int    `json:"CreateTime"`
	MsgType         string `json:"MsgType"`
	Event           string `json:"Event"`
	ProductSpuAudit struct {
		ProductID int    `json:"product_id"`
		Status    int    `json:"status"`
		Reason    string `json:"reason"`
	} `json:"ProductSpuAudit"`
}

func (EventProductSpuListing) GetMessageType() string {
	return "event"
}

func (EventProductSpuListing) GetEventType() string {
	return "product_spu_listing"
}

func (m EventProductSpuListing) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (EventProductSpuListing) ParseFromJson(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventProductSpuListing
	err := json.Unmarshal(data, &temp)
	return temp, err
}
