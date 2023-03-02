package callbacks

import "encoding/json"

// 结算账户变更回调
// 文档: https://developers.weixin.qq.com/doc/channels/API/funds/callback/channels_ec_acct_notify.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcAcctNotify{})
}

type ChannelsEcAcctNotify struct {
	CreateTime  int    `json:"CreateTime"`
	Event       string `json:"Event"`
	MsgType     string `json:"MsgType"`
	ToUserName  string `json:"ToUserName"`
	AccountInfo struct {
		Event int `json:"event"`
	} `json:"account_info"`
}

func (ChannelsEcAcctNotify) GetMessageType() string {
	return "event"
}

func (ChannelsEcAcctNotify) GetEventType() string {
	return "channels_ec_acct_notify"
}

func (m ChannelsEcAcctNotify) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcAcctNotify) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcAcctNotify
	err := json.Unmarshal(data, &temp)
	return temp, err
}
