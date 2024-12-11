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

// 商品编辑
const EventTypeProductSpuUpdate EventType = "product_spu_update"

// 优惠券领取通知
const EventTypeChannelsEcCouponReceive EventType = "channels_ec_coupon_receive"

// 类目审核结果
const EventTypeProductCategoryAudit EventType = "product_category_audit"

// 品牌资质事件回调
const EventTypeChannelsEcBrand EventType = "channels_ec_brand"

// 订单下单
const EventTypeChannelsEcOrderNew EventType = "channels_ec_order_new"

// 订单取消
const EventTypeChannelsEcOrderCancel EventType = "channels_ec_order_cancel"

// 订单支付成功
const EventTypeChannelsEcOrderPay EventType = "channels_ec_order_pay"

// 订单发货
const EventTypeChannelsEcOrderDeliver EventType = "channels_ec_order_deliver"

// 订单确认收货
const EventTypeChannelsEcOrderConfirm EventType = "channels_ec_order_confirm"

// 订单结算成功
const EventTypeChannelsEcOrderSettle EventType = "channels_ec_order_settle"

// 售后单更新通知
const EventTypeChannelsEcAftersaleUpdate EventType = "channels_ec_aftersale_update"

// 纠纷更新通知
const EventTypeChannelsEcComplaintUpdate EventType = "channels_ec_complaint_update"

// 结算账户变更回调
const EventTypeChannelsEcAcctNotify EventType = "channels_ec_acct_notify"

// 提现回调
const EventTypeChannelsEcWithdrawNotify EventType = "channels_ec_withdraw_notify"

// 提现二维码回调
const EventTypeQrcodeStatus EventType = "qrcode_status"

// 订单其他信息更新
const EventTypeChannelsEcOrderExtInfoUpdate EventType = "channels_ec_order_ext_info_update"

// 创建优惠券通知
const EventTypeChannelsEcCouponCreate EventType = "channels_ec_coupon_create"

// 删除优惠券通知
const EventTypeChannelsEcCouponDelete EventType = "channels_ec_coupon_delete"

// 优惠券过期通知
const EventTypeChannelsEcCouponExpire EventType = "channels_ec_coupon_expire"

// 优惠券信息更新通知
const EventTypeChannelsEcCouponInfoChange EventType = "channels_ec_coupon_info_change"

// 作废优惠券通知
const EventTypeChannelsEcCouponInvalid EventType = "channels_ec_coupon_invalid"

// 用户优惠券过期通知
const EventTypeChannelsEcUserCouponExpire EventType = "channels_ec_user_coupon_expire"

// 优惠券返还通知
const EventTypeChannelsEcUserCouponUnuse EventType = "channels_ec_user_coupon_unuse"

// 优惠券核销通知
const EventTypeChannelsEcUserCouponUse EventType = "channels_ec_user_coupon_use"

// 物流轨迹推送通知
const EventTypeEwaybillPushPath EventType = "ewaybill_push_path"
