package channels

// 订单状态
const (
    OrderStatusNoPay          int = 10  // 待付款
    OrderStatusStocking       int = 20  // 待发货
    OrderStatusPartSend       int = 21  // 部分发货
    OrderStatusSend           int = 30  // 已发货
    OrderStatusAfsCancel      int = 200 // 全部商品售后之后，订单取消
    OrderStatusCustomerCancel int = 250 // 未付款用户主动取消或超时未付款订单自动取消
    OrderStatusFinish         int = 100 // 完成
)

// 售后类型
const (
    AfsTypeRefund = "REFUND" // 退款
    AfsTypeReturn = "RETURN" // 退货退款
)

// 售后状态
const (
    AfterSaleStatusUserCanceld             = "USER_CANCELD"               // 用户取消申请
    AfterSaleStatusMerchantProcessing      = "MERCHANT_PROCESSING"        // 商家受理中
    AfterSaleStatusMerchantRejectRefund    = "MERCHANT_REJECT_REFUND"     // 商家拒绝退款
    AfterSaleStatusMerchantRejectReturn    = "MERCHANT_REJECT_RETURN"     // 商家拒绝退货退款
    AfterSaleStatusUserWaitReturn          = "USER_WAIT_RETURN"           // 待买家退货
    AfterSaleStatusReturnClosed            = "RETURN_CLOSED"              // 退货退款关闭
    AfterSaleStatusMerchantWaitReceipt     = "MERCHANT_WAIT_RECEIPT"      // 待商家收货
    AfterSaleStatusMerchantOverdueRefund   = "MERCHANT_OVERDUE_REFUND"    // 商家逾期未退款
    AfterSaleStatusMerchantRefundSuccess   = "MERCHANT_REFUND_SUCCESS"    // 退款完成
    AfterSaleStatusMerchantReturnSuccess   = "MERCHANT_RETURN_SUCCESS"    // 退货退款完成
    AfterSaleStatusPlatformRefunding       = "PLATFORM_REFUNDING"         // 平台退款中
    AfterSaleStatusPlatformRefundFail      = "PLATFORM_REFUND_FAIL"       // 平台退款失败
    AfterSaleStatusUserWaitConfirm         = "USER_WAIT_CONFIRM"          // 待用户确认
    AfterSaleStatusMerchantRefundRetryFail = "MERCHANT_REFUND_RETRY_FAIL" // 商家打款失败，客服关闭售后
    AfterSaleStatusMerchantFail            = "MERCHANT_FAIL"              // 售后关闭
)

// 发货方式
const (
    DeliverTypeSelf    int = 1 // 自寄快递发货
    DeliverTypeVirtual int = 3 // 虚拟商品无需物流发货
)
