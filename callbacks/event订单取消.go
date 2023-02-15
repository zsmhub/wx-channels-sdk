package callbacks

import "encoding/json"

// 订单取消
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_cancel.html

func init() {
    //添加可解析的回调事件
    supportCallback(ChannelsEcOrderCancel{})
}

type ChannelsEcOrderCancel struct {
    CreateTime   int    `json:"CreateTime"`
    Event        string `json:"Event"`
    FromUserName string `json:"FromUserName"`
    MsgType      string `json:"MsgType"`
    ToUserName   string `json:"ToUserName"`
    OrderInfo    struct {
        CancelType int `json:"cancel_type"`
        OrderID    int `json:"order_id"`
    } `json:"order_info"`
}

func (ChannelsEcOrderCancel) GetMessageType() string {
    return "event"
}

func (ChannelsEcOrderCancel) GetEventType() string {
    return "channels_ec_order_cancel"
}

func (m ChannelsEcOrderCancel) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderCancel) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
    var temp ChannelsEcOrderCancel
    err := json.Unmarshal(data, &temp)
    return temp, err
}
