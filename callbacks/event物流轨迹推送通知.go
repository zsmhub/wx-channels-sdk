package callbacks

import "encoding/json"

// 物流轨迹推送通知
// 文档: https://developers.weixin.qq.com/doc/store/API/ewaybill/push_path.html

func init() {
	//添加可解析的回调事件
	supportCallback(EwaybillPushPath{})
}

type EwaybillPushPath struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	WaybillInfo  struct {
		Desc            string `json:"desc"`
		EwaybillOrderID int    `json:"ewaybill_order_id"`
		Status          int    `json:"status"`
		UpdateTime      int    `json:"update_time"`
		WaybillID       string `json:"waybill_id"`
	} `json:"waybill_info"`
}

func (EwaybillPushPath) GetMessageType() string {
	return "event"
}

func (EwaybillPushPath) GetEventType() string {
	return "ewaybill_push_path"
}

func (m EwaybillPushPath) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (EwaybillPushPath) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp EwaybillPushPath
	err := json.Unmarshal(data, &temp)
	return temp, err
}
