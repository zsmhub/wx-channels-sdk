package callbacks

import "encoding/json"

// 类目审核结果
// 文档: https://developers.weixin.qq.com/doc/channels/API/category/callback/ProductCategoryAudit.html

func init() {
	//添加可解析的回调事件
	supportCallback(ProductCategoryAudit{})
}

type ProductCategoryAudit struct {
	CreateTime           int    `json:"CreateTime"`
	Event                string `json:"Event"`
	FromUserName         string `json:"FromUserName"`
	MsgType              string `json:"MsgType"`
	ProductCategoryAudit struct {
		AuditID string `json:"audit_id"`
		Reason  string `json:"reason"`
		Status  string `json:"status"`
	} `json:"ProductCategoryAudit"`
	ToUserName string `json:"ToUserName"`
}

func (ProductCategoryAudit) GetMessageType() string {
	return "event"
}

func (ProductCategoryAudit) GetEventType() string {
	return "product_category_audit"
}

func (m ProductCategoryAudit) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ProductCategoryAudit) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ProductCategoryAudit
	err := json.Unmarshal(data, &temp)
	return temp, err
}
