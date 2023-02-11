package callbacks

import (
	"encoding/json"
)

// 商品更新通知
// 文档: 暂无

func init() {
	supportCallback(EventProductSpuUpdate{})
}

type EventProductSpuUpdate struct {
	ToUserName       string `json:"ToUserName"`
	FromUserName     string `json:"FromUserName"`
	CreateTime       int    `json:"CreateTime"`
	MsgType          string `json:"MsgType"`
	Event            string `json:"Event"`
	ProductSpuUpdate struct {
		ProductId int64 `json:"product_id"`
		Status    int   `json:"status"`
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

func (EventProductSpuUpdate) ParseFromJson(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventProductSpuUpdate
	err := json.Unmarshal(data, &temp)
	return temp, err
}
