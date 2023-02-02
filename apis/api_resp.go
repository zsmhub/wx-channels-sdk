package apis

type CommonResp struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
}

// IsOK 响应体是否为一次成功请求的响应
// > 开发者需根据errcode是否为0判断是否调用成功(errcode意义请见全局错误码)。
// > 而errmsg仅作参考，后续可能会有变动，因此不可作为是否调用成功的判据。
func (x *CommonResp) IsOK() bool {
	return x.ErrCode == 0
}

func (x *CommonResp) TryIntoErr() error {
	if x.IsOK() {
		return nil
	}

	return &ClientError{
		Code: x.ErrCode,
		Msg:  x.ErrMsg,
	}
}
