package callbacks

import (
	"encoding/json"
)

// 商品审核
// 文档: https://developers.weixin.qq.com/doc/channels/API/product/callback/ProductSpuAudit.html

func init() {
	supportCallback(EventProductSpuAudit{})
}

type EventProductSpuAudit struct {
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

func (EventProductSpuAudit) GetMessageType() string {
	return "event"
}

func (EventProductSpuAudit) GetEventType() string {
	return "product_spu_audit"
}

func (m EventProductSpuAudit) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (EventProductSpuAudit) ParseFromJson(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventProductSpuAudit
	err := json.Unmarshal(data, &temp)
	return temp, err
}
