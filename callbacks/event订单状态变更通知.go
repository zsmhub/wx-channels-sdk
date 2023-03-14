package callbacks

import (
	"github.com/tidwall/gjson"
)

// 订单状态变更通知【注意该回调跟其他订单状态回调重复了，如下单，会同时推“送订单下单”和“订单状态变更通知”两个回调事件】
// 文档: 暂无

func init() {
	//添加可解析的回调事件
	supportCallback(ProductOrderStatusUpdate{})
}

type ProductOrderStatusUpdate struct {
	CreateTime               int64  `json:"CreateTime"`
	Event                    string `json:"Event"`
	FromUserName             string `json:"FromUserName"`
	MsgType                  string `json:"MsgType"`
	ToUserName               string `json:"ToUserName"`
	ProductOrderStatusUpdate struct {
		OrderID string `json:"order_id"`
		Status  int64  `json:"status"`
	} `json:"ProductOrderStatusUpdate"`
}

func (ProductOrderStatusUpdate) GetMessageType() string {
	return "event"
}

func (ProductOrderStatusUpdate) GetEventType() string {
	return "product_order_status_update"
}

func (m ProductOrderStatusUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ProductOrderStatusUpdate) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ProductOrderStatusUpdate{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		ProductOrderStatusUpdate: struct {
			OrderID string `json:"order_id"`
			Status  int64  `json:"status"`
		}{
			OrderID: gjson.GetBytes(data, "order_info.order_id").String(),
			Status:  gjson.GetBytes(data, "order_info.status").Int(),
		},
	}
	return temp, nil
}
