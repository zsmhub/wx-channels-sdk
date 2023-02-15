package callbacks

import "encoding/json"

// 订单发货
// 文档: https://developers.weixin.qq.com/doc/channels/API/order/callback/channels_ec_order_deliver.html

func init() {
    //添加可解析的回调事件
    supportCallback(ChannelsEcOrderDeliver{})
}

type ChannelsEcOrderDeliver struct {
    CreateTime   int    `json:"CreateTime"`
    Event        string `json:"Event"`
    FromUserName string `json:"FromUserName"`
    MsgType      string `json:"MsgType"`
    ToUserName   string `json:"ToUserName"`
    OrderInfo    struct {
        FinishDelivery int `json:"finish_delivery"`
        OrderID        int `json:"order_id"`
    } `json:"order_info"`
}

func (ChannelsEcOrderDeliver) GetMessageType() string {
    return "event"
}

func (ChannelsEcOrderDeliver) GetEventType() string {
    return "channels_ec_order_deliver"
}

func (m ChannelsEcOrderDeliver) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderDeliver) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
    var temp ChannelsEcOrderDeliver
    err := json.Unmarshal(data, &temp)
    return temp, err
}
