package callbacks

import "github.com/tidwall/gjson"

// 商品上下架
// 文档: https://developers.weixin.qq.com/doc/channels/API/product/callback/ProductSpuListing.html

func init() {
	supportCallback(EventProductSpuListing{})
}

type EventProductSpuListing struct {
	CreateTime        int64  `json:"CreateTime"`
	Event             string `json:"Event"`
	FromUserName      string `json:"FromUserName"`
	MsgType           string `json:"MsgType"`
	ToUserName        string `json:"ToUserName"`
	ProductSpuListing struct {
		ProductID string `json:"product_id"`
		Status    int64  `json:"status"`
		Reason    string `json:"reason"`
	} `json:"ProductSpuListing"`
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

func (EventProductSpuListing) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = EventProductSpuListing{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		ProductSpuListing: struct {
			ProductID string `json:"product_id"`
			Status    int64  `json:"status"`
			Reason    string `json:"reason"`
		}{
			ProductID: gjson.GetBytes(data, "ProductSpuListing.product_id").String(),
			Status:    gjson.GetBytes(data, "ProductSpuListing.status").Int(),
			Reason:    gjson.GetBytes(data, "ProductSpuListing.reason").String(),
		},
	}
	return temp, nil
}
