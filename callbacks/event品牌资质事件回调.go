package callbacks

import "encoding/json"

// 品牌资质事件回调
// 文档: https://developers.weixin.qq.com/doc/channels/API/brand/callback/brand_event.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcBrand{})
}

type ChannelsEcBrand struct {
	BrandEvent struct {
		AuditID string `json:"audit_id"`
		BrandID string `json:"brand_id"`
		Reason  string `json:"reason"`
		Status  int    `json:"status"`
	} `json:"BrandEvent"`
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
}

func (ChannelsEcBrand) GetMessageType() string {
	return "event"
}

func (ChannelsEcBrand) GetEventType() string {
	return "channels_ec_brand"
}

func (m ChannelsEcBrand) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcBrand) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp ChannelsEcBrand
	err := json.Unmarshal(data, &temp)
	return temp, err
}
