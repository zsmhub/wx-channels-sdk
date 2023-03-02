package callbacks

import "encoding/json"

// 提现二维码回调
// 文档: https://developers.weixin.qq.com/doc/channels/API/funds/callback/qrcode_status.html

func init() {
	//添加可解析的回调事件
	supportCallback(QrcodeStatus{})
}

type QrcodeStatus struct {
	CreateTime int    `json:"CreateTime"`
	Event      string `json:"Event"`
	MsgType    string `json:"MsgType"`
	ToUserName string `json:"ToUserName"`
	QrcodeInfo struct {
		ScanUserType int    `json:"scan_user_type"`
		Status       int    `json:"status"`
		Ticket       string `json:"ticket"`
	} `json:"qrcode_info"`
}

func (QrcodeStatus) GetMessageType() string {
	return "event"
}

func (QrcodeStatus) GetEventType() string {
	return "qrcode_status"
}

func (m QrcodeStatus) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (QrcodeStatus) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp QrcodeStatus
	err := json.Unmarshal(data, &temp)
	return temp, err
}
