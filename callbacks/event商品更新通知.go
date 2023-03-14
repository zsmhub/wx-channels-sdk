package callbacks

import "github.com/tidwall/gjson"

// 商品更新通知
// 文档: 暂无

func init() {
	supportCallback(EventProductSpuUpdate{})
}

type EventProductSpuUpdate struct {
	CreateTime       int64  `json:"CreateTime"`
	Event            string `json:"Event"`
	FromUserName     string `json:"FromUserName"`
	MsgType          string `json:"MsgType"`
	ToUserName       string `json:"ToUserName"`
	ProductSpuUpdate struct {
		ProductID string `json:"product_id"`
		Status    int64  `json:"status"`
	} `json:"ProductSpuUpdate"`
}

func (EventProductSpuUpdate) GetMessageType() string {
	return "event"
}

func (EventProductSpuUpdate) GetEventType() string {
	return "product_spu_update"
}

func (m EventProductSpuUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (EventProductSpuUpdate) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = EventProductSpuUpdate{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		ProductSpuUpdate: struct {
			ProductID string `json:"product_id"`
			Status    int64  `json:"status"`
		}{
			ProductID: gjson.GetBytes(data, "ProductSpuUpdate.product_id").String(),
			Status:    gjson.GetBytes(data, "ProductSpuUpdate.status").Int(),
		},
	}
	return temp, nil
}
