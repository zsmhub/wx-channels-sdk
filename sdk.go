package channels

import (
	"errors"
	"github.com/zsmhub/wx-channels-sdk/apis"
	"github.com/zsmhub/wx-channels-sdk/callbacks"
)

// 微信视频号 sdk 调用入口
var Sdk = &sdk{}

type sdk struct {
	// 视频号小店
	ShopCallback *callbacks.CallbackHandler // 回调事件解析
	ShopClient   *apis.ApiClient            // API客户端，用于视频号小店的接口调用

	// 视频号橱窗
	WindowCallback *callbacks.CallbackHandler // 回调事件解析
	WindowClient   *apis.ApiClient            // API客户端，用于视频号橱窗的接口调用

	apis.Options
}

// 初始化sdk参数
func (s *sdk) InitOptions(opts apis.Options) {
	s.Options = opts
}

// 视频号小店：回调事件初始化 (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewShopCallbackHandler(token, encodingAESKey string) error {
	if token == "" || encodingAESKey == "" {
		return errors.New("传参不能为空")
	}

	var err error
	s.ShopCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)

	return err
}

// 视频号小店：API客户端初始化
func (s *sdk) NewShopApiClient(appId, appSecret string) {
	s.ShopClient = apis.NewApiClient(appId, appSecret, s.Options)
}

// 视频号橱窗：回调事件初始化 (如果要处理微信回调, 这一步是必须的)
func (s *sdk) NewWindowCallbackHandler(token, encodingAESKey string) error {
	if token == "" || encodingAESKey == "" {
		return errors.New("传参不能为空")
	}

	var err error
	s.WindowCallback, err = callbacks.NewCallbackHandler(token, encodingAESKey)

	return err
}

// 视频号橱窗：API客户端初始化
func (s *sdk) NewWindowApiClient(appId, appSecret string) {
	s.WindowClient = apis.NewApiClient(appId, appSecret, s.Options)
}
