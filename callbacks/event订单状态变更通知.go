package callbacks

import "encoding/json"

// 订单状态变更通知【注意该回调跟其他订单状态回调重复了，如下单，会同时推“送订单下单”和“订单状态变更通知”两个回调事件】
// 文档: 暂无

func init() {
	//添加可解析的回调事件
	supportCallback(ProductOrderStatusUpdate{})
}

type ProductOrderStatusUpdate struct {
	ToUserName               string `json:"ToUserName"`
	FromUserName             string `json:"FromUserName"`
	CreateTime               int    `json:"CreateTime"`
	MsgType                  string `json:"MsgType"`
	Event                    string `json:"Event"`
	ProductOrderStatusUpdate struct {
		OrderId int64 `json:"order_id"`
		Status  int   `json:"status"`
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
	var temp ProductOrderStatusUpdate
	err := json.Unmarshal(data, &temp)
	return temp, err
}
