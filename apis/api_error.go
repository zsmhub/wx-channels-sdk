package apis

import "fmt"

// ClientError 响应错误
type ClientError struct {
	// Code 错误码，0表示请求成功，非0表示请求失败。
	// 开发者需根据errcode是否为0判断是否调用成功。
	Code ErrCode
	// Msg 错误信息，调用失败会有相关的错误信息返回。
	// 仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
	Msg string
}

var _ error = (*ClientError)(nil)

func (e *ClientError) Error() string {
	return fmt.Sprintf(
		"ClientError { Code: %d, Msg: %#v }",
		e.Code,
		e.Msg,
	)
}

// ErrCode 错误码类型，目前只能在各个接口文档下查看可能出现的错误码类型
type ErrCode int64

// 请求的参数错误
const ErrCodeMinus2 ErrCode = -2

// 系统异常或系统繁忙，请开发者稍候再试
const ErrCodeSysErr ErrCode = -1

// 请求成功
const ErrCodeSuccess ErrCode = 0

// AppSecret 错误或者 AppSecret 不属于这个小店，请开发者确认 AppSecret 的正确性
const ErrCode40001 ErrCode = 40001

// 请确保 grant_type 字段值为 client_credential
const ErrCode40002 ErrCode = 40002

// 不合法的 AppID，请开发者检查 AppID 的正确性，避免异常字符，注意大小写
const ErrCode40013 ErrCode = 40013

// 需要使用 HTTP GET
const ErrCode43001 ErrCode = 43001

// 无效的地址编码
const ErrCode10021402 ErrCode = 10021402

// 参数有误
const ErrCode10020055 ErrCode = 10020055

// 无权限调用该api
const ErrCode10020050 ErrCode = 10020050

// 商品不存在
const ErrCode10020052 ErrCode = 10020052

// 商品未上架
const ErrCode10020065 ErrCode = 10020065

// 参数有误
const ErrCode9401020 ErrCode = 9401020

// 资质材料超限，最多不能超过10张
const ErrCode10020062 ErrCode = 10020062

// 不合法的类目ID
const ErrCode10020063 ErrCode = 10020063

// 该类目审核中，无需重复提交
const ErrCode10020064 ErrCode = 10020064

// 不合法的fileid
const ErrCode10020087 ErrCode = 10020087

// 必须使用上传资质接口获取fileid
const ErrCode10020094 ErrCode = 10020094

// 不合法的审核ID
const ErrCode10020006 ErrCode = 10020006

// 提交的资质数量超出限制
const ErrCode10020080 ErrCode = 10020080

// 品牌重复申请
const ErrCode10020082 ErrCode = 10020082

// 品牌名称与品牌库 ID 不匹配
const ErrCode10020084 ErrCode = 10020084

// 品牌资质文件的file_id无效
const ErrCode10020086 ErrCode = 10020086

// 品牌的申请记录不存在，请调用新增品牌接口
const ErrCode10020081 ErrCode = 10020081

// 该品牌申请正在审核中，请不要重复提交
const ErrCode10020085 ErrCode = 10020085

// 当前商品不允许编辑
const ErrCode10020008 ErrCode = 10020008

// 商品的类目长度不对（预期是有三级类目）
const ErrCode10020011 ErrCode = 10020011

// 销售属性不合法，不属于商品所属的三级类目
const ErrCode10020012 ErrCode = 10020012

// 商品 sku 数量不合理
const ErrCode10020013 ErrCode = 10020013

// 批量添加 sku 失败
const ErrCode10020016 ErrCode = 10020016

// 类目非法
const ErrCode10020017 ErrCode = 10020017

// 商家不具备当前类目资质
const ErrCode10020018 ErrCode = 10020018

// 运费模版非法
const ErrCode10020019 ErrCode = 10020019

// 商品标题为空
const ErrCode10020020 ErrCode = 10020020

// 商品标题过长
const ErrCode10020021 ErrCode = 10020021

// 商品头图为空
const ErrCode10020022 ErrCode = 10020022

// 商品头图过多
const ErrCode10020023 ErrCode = 10020023

// 商品描述过长
const ErrCode10020024 ErrCode = 10020024

// 商品详情图片过多
const ErrCode10020025 ErrCode = 10020025

// 商品详情描述过长
const ErrCode10020026 ErrCode = 10020026

// 资质图片过多
const ErrCode10020027 ErrCode = 10020027

// sku价格过高
const ErrCode10020028 ErrCode = 10020028

// sku商品编码过长
const ErrCode10020029 ErrCode = 10020029

// sku_out_id已存在
const ErrCode10020030 ErrCode = 10020030

// sku销售属性相同 key 下不能超过100个不同value
const ErrCode10020031 ErrCode = 10020031

// sku销售属性 key 过长
const ErrCode10020032 ErrCode = 10020032

// sku销售属性 value 过长
const ErrCode10020033 ErrCode = 10020033

// 图片/视频 url 非法， url 前缀应为mmecimage.cn/p/
const ErrCode10020035 ErrCode = 10020035

// out_product_id过长
const ErrCode10020036 ErrCode = 10020036

// out_sku_id过长
const ErrCode10020037 ErrCode = 10020037

// 上架的商品缺少sku
const ErrCode10020038 ErrCode = 10020038

// SKU价格为0
const ErrCode10020039 ErrCode = 10020039

// sku售卖价格大于市场价格
const ErrCode10020040 ErrCode = 10020040

// 账号注销中
const ErrCode10020041 ErrCode = 10020041

// 商品标题过短
const ErrCode10020042 ErrCode = 10020042

// 类目不可用，请更换类目
const ErrCode10020043 ErrCode = 10020043

// 商品标题不得仅为数字、字母、字符，不得含非法字符，请修改后重新提交
const ErrCode10020045 ErrCode = 10020045

// 商品信息设置有误，请重新输入
const ErrCode10020046 ErrCode = 10020046

// 由于类目保证金不足，已禁止新增商品
const ErrCode10020048 ErrCode = 10020048

// 参数错误
const ErrCode10020051 ErrCode = 10020051

// 提审频率达到限制
const ErrCode10020066 ErrCode = 10020066

// 当前运费模版计价方式为[按重量]，且提交商品重量为0
const ErrCode10020068 ErrCode = 10020068

// 当前类目不允许选择无需快递的发货方式
const ErrCode10020069 ErrCode = 10020069

// 由于类目保证金不足，已下架所有商品
const ErrCode10020070 ErrCode = 10020070

// 该商品所需类目保证金高于当前保证金余额，请前往商家网页端（https://channels.weixin.qq.com/shop）添加一次该类目商品，即可完成保证金补缴。
const ErrCode10020083 ErrCode = 10020083

// 当前类目不支持当前品牌，或商品品牌 id 非法，或商品品牌 id 未申请通过
const ErrCode10020088 ErrCode = 10020088

// 售后说明超过长度限制（200 UTF字符）
const ErrCode10020095 ErrCode = 10020095

// 商品参数重复
const ErrCode10020096 ErrCode = 10020096

// 商品参数缺少必填项
const ErrCode10020097 ErrCode = 10020097

// 商品参数内容有误
const ErrCode10020098 ErrCode = 10020098

// 商品正在审核中，无法编辑或删除，请先调用撤回商品审核接口
const ErrCode10020049 ErrCode = 10020049

// 该 spu 处于抢购状态，不能修改sku
const ErrCode10020014 ErrCode = 10020014

// 该商品的sku_id/sku_out_id填写有误
const ErrCode10020034 ErrCode = 10020034

// 该商品上一次提交正在上传中，请稍后再试
const ErrCode10020067 ErrCode = 10020067

// 未开张商店发布商品
const ErrCode10020089 ErrCode = 10020089

// 暂不支持上架该类型的商品
const ErrCode10020091 ErrCode = 10020091

// 商品edit_status != 2
const ErrCode10020044 ErrCode = 10020044

// 当前商品状态不允许直接更新库存，或product_id和sku_id不匹配，或库存不足以减扣，或增加减少数量为0
const ErrCode10020054 ErrCode = 10020054

// 抢购商品未上架
const ErrCode10020071 ErrCode = 10020071

// 抢购开始/结束时间无效，请设置一年内的时间
const ErrCode10020072 ErrCode = 10020072

// 抢购库存超出商品可用库存
const ErrCode10020073 ErrCode = 10020073

// 设置抢购的 SKU 无效
const ErrCode10020074 ErrCode = 10020074

// 商品已存在其他抢购活动
const ErrCode10020075 ErrCode = 10020075

// 抢购价格超出原价
const ErrCode10020078 ErrCode = 10020078

// 接口暂时下线
const ErrCode10020090 ErrCode = 10020090

// 抢购活动不存在
const ErrCode10020076 ErrCode = 10020076

// out_warehouse_id命名非法
const ErrCode10020200 ErrCode = 10020200

// address_id非法，或关联关系非法
const ErrCode10020203 ErrCode = 10020203

// 仓库名称或仓库介绍为空
const ErrCode10020205 ErrCode = 10020205

// out_warehouse_id不存在
const ErrCode10020202 ErrCode = 10020202

// 重复的out_warehouse_id
const ErrCode10020206 ErrCode = 10020206

// 该 sku 设置区域仓数量超过限制
const ErrCode10020204 ErrCode = 10020204

// 请求体格式不正确，请检查请求体中各个参数的类型是否正确
const ErrCode47001 ErrCode = 47001

// 请求体参数不正确，请检查各个参数是否按规范填写，具体原因查看errmsg
const ErrCode40097 ErrCode = 40097

// 无权调用本api，请检查相关权限是否已开通
const ErrCode48001 ErrCode = 48001

// 订单不存在，请检查订单号与 token 是否正确
const ErrCode100002 ErrCode = 100002

// 订单不是未付款订单，不能改价
const ErrCode101100 ErrCode = 101100

// 该订单改价次数已经超过50次
const ErrCode108022 ErrCode = 108022

// 某种商品或者运费修改后的价格超出原价
const ErrCode108021 ErrCode = 108021

// 修改后订单价格为0
const ErrCode108023 ErrCode = 108023

// 订单状态不正确，待发货订单才可以修改地址
const ErrCode606000 ErrCode = 606000

// 修改地址次数已经达到上限，不可超过5次
const ErrCode606001 ErrCode = 606001

// user_address字段未经过检查，请检查是否有必填字段未填写
const ErrCode606002 ErrCode = 606002

// 售后单号不存在
const ErrCode10020000 ErrCode = 10020000

// 售后单状态不合法
const ErrCode10020001 ErrCode = 10020001

// 当前售后单不支持此操作，请检查售后单状态
const ErrCode10021041 ErrCode = 10021041

// 售后单非当前账号，请检查参数
const ErrCode10021043 ErrCode = 10021043

// 售后单不存在
const ErrCode10021050 ErrCode = 10021050

// 非法的时间区间，不得大于24小时
const ErrCode10020004 ErrCode = 10020004

// 当前用户太多，请稍后再试
const ErrCode10021044 ErrCode = 10021044

// 退款失败
const ErrCode10021045 ErrCode = 10021045

// 金额不足
const ErrCode10021046 ErrCode = 10021046

// 纠纷单号不存在
const ErrCode10021500 ErrCode = 10021500

// 单条留言文字数量过多
const ErrCode10021501 ErrCode = 10021501

// 所有留言图片总数过多
const ErrCode10021502 ErrCode = 10021502

// 不是待客服处理的状态，无法留言
const ErrCode10021503 ErrCode = 10021503

// 不是待商家举证的状态，无法留言
const ErrCode10021504 ErrCode = 10021504

// 订单更新过于频繁，更新订单失败，请重试
const ErrCode100003 ErrCode = 100003

// 当前订单有未完成售后单，不允许发货
const ErrCode108009 ErrCode = 108009

// 发货失败，请检查 errmsg 中返回的商品是否正确填写，包括订单中是否包含该商品，该商品是否已经发货，是否已经完成售后无法发货，以及商品数量是否填写正确等
const ErrCode109000 ErrCode = 109000

// 当前订单已经发货完成，不能重复发货
const ErrCode109001 ErrCode = 109001

// 当前发货请求里没有带上发货商品，请检查product_infos和is_all_product是否按规范填写
const ErrCode109002 ErrCode = 109002

// 当前订单已经下单在线物流单，不允许自寄发货
const ErrCode109205 ErrCode = 109205

// 商家没有默认退货地址，不可发货，请设置默认退货地址后再进行发货
const ErrCode606003 ErrCode = 606003

// delivery_id不合法，请使用【获取物流公司列表】接口获取合法的delivery_id
const ErrCode606004 ErrCode = 606004

// waybill_id不合法，请检查快递单号是否正确填写
const ErrCode606005 ErrCode = 606005

// deliver_type不符合要求，普通订单请使用快递发货，虚拟商品订单请无需物流发货，请检查订单deliver_method
const ErrCode606008 ErrCode = 606008

// 同一 sku 的所有商品必须在一个包裹里发出
const ErrCode606030 ErrCode = 606030

// 发货请求里product_id或sku_id填写不正确，请检查订单中是否包含该商品
const ErrCode606031 ErrCode = 606031

// 当前请求方不能对该订单进行发货，请使用特定第三方服务商身份进行发货操作
const ErrCode606034 ErrCode = 606034

// 非法的默认运费模板规则，默认运费模板不可以指定地址
const ErrCode10021053 ErrCode = 10021053

// 非法的省市区信息
const ErrCode10021054 ErrCode = 10021054

// 请检查模板名是否已存在｜
const ErrCode10021200 ErrCode = 10021200

// 验证视频号身份失败, 请检查是否使用视频号橱窗 ID 请求
const ErrCode10022001 ErrCode = 10022001

// 因违规行为, 橱窗被禁止使用, 请前往'带货中心->个人中心->带货权限'检查橱窗带货权限
const ErrCode10022002 ErrCode = 10022002

// 橱窗商品未找到
const ErrCode10022003 ErrCode = 10022003

// 商品异常，禁止添加到橱窗，请检查商品的banned_detail字段
const ErrCode10022004 ErrCode = 10022004

// 橱窗中的商品数超过1万，请减少商品数目后再重试
const ErrCode10022005 ErrCode = 10022005

// 不支持操作带货中心来源的商品
const ErrCode10022007 ErrCode = 10022007

// 请求的 appid 不属于该视频号的绑定店铺
const ErrCode10022008 ErrCode = 10022008

// 请求参数错误，请检查请求参数是否符合限制
const ErrCode10022006 ErrCode = 10022006

// 请求的request_id不存在
const ErrCode1 ErrCode = 1

// 没有权限访问当前request_id对应的留资数据
const ErrCode100000 ErrCode = 100000

// 图片格式不合法, 只支持 bmp, jpeg, jpg, png, gif, svg
const ErrCode10020056 ErrCode = 10020056

// 获取原图失败，请检查 url 是否合法
const ErrCode10020057 ErrCode = 10020057

// 上传图片失败，请重试
const ErrCode10020058 ErrCode = 10020058

// 图片为空
const ErrCode10020059 ErrCode = 10020059

// 图片大小超出限制
const ErrCode10020060 ErrCode = 10020060

// 图片 URL 非法，如：传入mmecimage.cn/p/前缀的图片
const ErrCode10020061 ErrCode = 10020061

// 不合法的 url ，递交的页面被 sitemap 标记为拦截
const ErrCode40066 ErrCode = 40066

// 优惠券名称太长
const ErrCode10021005 ErrCode = 10021005

// 校验折扣数失败
const ErrCode10021006 ErrCode = 10021006

// 校验优惠价格失败
const ErrCode10021007 ErrCode = 10021007

// 校验直减券是否小于最低价格
const ErrCode10021008 ErrCode = 10021008

// 校验领取时间失败
const ErrCode10021009 ErrCode = 10021009

// 校验有效时间失败
const ErrCode10021010 ErrCode = 10021010

// 校验优惠券总发放量失败
const ErrCode10021011 ErrCode = 10021011

// 校验限领失败
const ErrCode10021012 ErrCode = 10021012

// 校验商户失败
const ErrCode10021013 ErrCode = 10021013

// 推广类型不对
const ErrCode10021014 ErrCode = 10021014

// 校验入参失败，含有非商户的指定商品
const ErrCode10021021 ErrCode = 10021021

// 优惠券信息违规
const ErrCode10021024 ErrCode = 10021024

// 创建 优惠券类型 暂不支持
const ErrCode10021035 ErrCode = 10021035

// 优惠券不存在
const ErrCode10021037 ErrCode = 10021037

// 优惠券状态不对
const ErrCode10021001 ErrCode = 10021001

// 优惠券还没过期
const ErrCode10021003 ErrCode = 10021003

// 优惠券重入失败
const ErrCode10021034 ErrCode = 10021034

// 分页拉优惠券列表，页码与上次请求相差过大
const ErrCode10021036 ErrCode = 10021036

// 暂无数据
const ErrCode10021302 ErrCode = 10021302

// 暂无数据
const ErrCode9710001 ErrCode = 9710001

// ticket已失效
const ErrCode60220 ErrCode = 60220

// 错误的ticket
const ErrCode60208 ErrCode = 60208

// 电子面单帐号列表为空，请检查请求参数
const ErrCode10025120 ErrCode = 10025120

// 快递公司不支持
const ErrCode10025102 ErrCode = 10025102

// delivery_id错误
const ErrCode10025012 ErrCode = 10025012

// 网点信息错误
const ErrCode10025013 ErrCode = 10025013

// 网点账号编码错误
const ErrCode10025014 ErrCode = 10025014

// 寄件人信息错误
const ErrCode10025015 ErrCode = 10025015

// 收件人信息错误
const ErrCode10025016 ErrCode = 10025016

// 小店订单信息错误
const ErrCode10025017 ErrCode = 10025017

// 店铺id信息错误
const ErrCode10025018 ErrCode = 10025018

// 下单失败，具体原因参考delivery_error_msg字段
const ErrCode10025005 ErrCode = 10025005

// 面单已存在
const ErrCode10025019 ErrCode = 10025019

// 取消下单失败，具体原因参考delivery_error_msg字段
const ErrCode10025006 ErrCode = 10025006

// 电子面单不存在
const ErrCode10025010 ErrCode = 10025010

// 更新冲突，请稍后重试
const ErrCode10025011 ErrCode = 10025011
