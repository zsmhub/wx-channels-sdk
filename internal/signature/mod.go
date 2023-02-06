package signature

import (
	"crypto/sha1" // nolint: gosec  // this is part of vendor API spec
	"crypto/subtle"
	"fmt"
	"net/url"
	"sort"
)

func MakeDevMsgSignature(paramValues ...string) string {
	tmp := make([]string, len(paramValues))
	copy(tmp, paramValues)

	sort.Strings(tmp)

	// nolint: gosec
	// this is part of vendor API spec
	state := sha1.New()
	for _, x := range tmp {
		_, _ = state.Write([]byte(x))
	}

	result := state.Sum(nil)
	return fmt.Sprintf("%x", result)
}

// ToSignature 适配企业微信请求参数签名的 interface
type ToSignature interface {
	// GetSignature 取请求上携带的签名串
	GetSignature() (string, bool)
	// GetParamValues 取所需请求参数值（不必有序）
	GetParamValues() ([]string, bool)
}

// VerifySignature 校验签名是否完好
func VerifySignature(token string, x ToSignature) bool {
	signature, ok := x.GetSignature()
	if !ok {
		return false
	}

	paramValues, ok := x.GetParamValues()
	if !ok {
		return false
	}

	devMsgSignature := MakeDevMsgSignature(append(paramValues, token)...)
	eq := subtle.ConstantTimeCompare([]byte(signature), []byte(devMsgSignature))
	return eq != 0
}

// VerifyHTTPRequestSignature 校验一个 HTTP 请求的签名是否完好
//
// 这是 VerifySignature 的简单包装。
func VerifyHTTPRequestSignature(token string, url *url.URL) bool {
	wrapped := httpRequestWithSignature{
		url: url,
	}
	return VerifySignature(token, &wrapped)
}

// httpRequestWithSignature 为 HTTP 请求适配签名校验逻辑
type httpRequestWithSignature struct {
	url *url.URL
}

var _ ToSignature = (*httpRequestWithSignature)(nil)

// GetSignature 取请求上携带的签名串
func (u *httpRequestWithSignature) GetSignature() (string, bool) {
	l := u.url.Query()["signature"]
	if len(l) != 1 {
		return "", false
	}

	return l[0], true
}

// GetParamValues 取所需请求参数值（不必有序）
func (u *httpRequestWithSignature) GetParamValues() ([]string, bool) {
	result := make([]string, 0)
	for k, l := range u.url.Query() {
		if k == "timestamp" || k == "nonce" {
			result = append(result, l...)
		}
	}
	return result, true
}
