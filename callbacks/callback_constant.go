package callbacks

// 回调消息/事件类型
// - MsgType: 消息类型，unknown(自定义的未知事件)，text/image/voice等消息事件， event(其他事件)
// - EventType: 事件类型，当 MsgType 为 event 时存在

// 消息类型
type MessageType string

// 文本消息
const MessageTypeText MessageType = "text"

// 图片消息
const MessageTypeImage MessageType = "image"

// 语音消息
const MessageTypeVoice MessageType = "voice"

// 视频消息
const MessageTypeVideo MessageType = "video"

// 位置消息
const MessageTypeLocation MessageType = "location"

// 链接消息
const MessageTypeLink MessageType = "link"

// 事件消息
const MessageTypeEvent MessageType = "event"

// 没有明确指定消息类型，都视为未知事件
const MessageTypeUnknown MessageType = "unknown"

// 事件类型
type EventType string

// 商品上下架
const EventTypeProductSpuListing EventType = "product_spu_listing"

// 商品审核
const EventTypeProductSpuAudit EventType = "product_spu_audit"
