package callbacks

import "encoding/json"

// 订单结算成功
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_settle.html

func init() {
    //添加可解析的回调事件
    supportCallback(ChannelsEcOrderSettle{})
}

type ChannelsEcOrderSettle struct {
    CreateTime   int    `json:"CreateTime"`
    Event        string `json:"Event"`
    FromUserName string `json:"FromUserName"`
    MsgType      string `json:"MsgType"`
    ToUserName   string `json:"ToUserName"`
    OrderInfo    struct {
        OrderID    int `json:"order_id"`
        SettleTime int `json:"settle_time"`
    } `json:"order_info"`
}

func (ChannelsEcOrderSettle) GetMessageType() string {
    return "event"
}

func (ChannelsEcOrderSettle) GetEventType() string {
    return "channels_ec_order_settle"
}

func (m ChannelsEcOrderSettle) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderSettle) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
    var temp ChannelsEcOrderSettle
    err := json.Unmarshal(data, &temp)
    return temp, err
}
