package callbacks

import "encoding/json"

// 提现回调
// 文档: https://developers.weixin.qq.com/doc/channels/API/funds/callback/channels_ec_withdraw_notify.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcWithdrawNotify{})
}

type ChannelsEcWithdrawNotify struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	WithdrawInfo struct {
		Event      int    `json:"event"`
		WithdrawID string `json:"withdraw_id"`
	} `json:"withdraw_info"`
}

func (ChannelsEcWithdrawNotify) GetMessageType() string {
	return "event"
}

func (ChannelsEcWithdrawNotify) GetEventType() string {
	return "channels_ec_withdraw_notify"
}

func (m ChannelsEcWithdrawNotify) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcWithdrawNotify) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcWithdrawNotify
	err := json.Unmarshal(data, &temp)
	return temp, err
}
