package callbacks

import "github.com/tidwall/gjson"

// 商品审核
// 文档: https://developers.weixin.qq.com/doc/channels/API/product/callback/ProductSpuAudit.html

func init() {
	supportCallback(EventProductSpuAudit{})
}

type EventProductSpuAudit struct {
	CreateTime      int64  `json:"CreateTime"`
	Event           string `json:"Event"`
	FromUserName    string `json:"FromUserName"`
	MsgType         string `json:"MsgType"`
	ToUserName      string `json:"ToUserName"`
	ProductSpuAudit struct {
		ProductID string `json:"product_id"`
		Status    int64  `json:"status"`
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

func (EventProductSpuAudit) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = EventProductSpuAudit{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		ProductSpuAudit: struct {
			ProductID string `json:"product_id"`
			Status    int64  `json:"status"`
			Reason    string `json:"reason"`
		}{
			ProductID: gjson.GetBytes(data, "ProductSpuAudit.product_id").String(),
			Status:    gjson.GetBytes(data, "ProductSpuAudit.status").Int(),
			Reason:    gjson.GetBytes(data, "ProductSpuAudit.reason").String(),
		},
	}
	return temp, nil
}
