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

// 系统繁忙，请开发者稍候再试
const ErrCodeServiceUnavailable ErrCode = -1

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
