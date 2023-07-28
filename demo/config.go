package demo

// 配置文件

// 视频号小店配置，可在「视频号小店 - 服务市场 - 自研」中获得。
var (
	ShopAppId                  = "xxx"
	ShopAppSecret              = "xxx"
	ShopCallbackToken          = "xxx"
	ShopCallbackEncodingAESKey = "xxx"
)

// 视频号橱窗配置，可在「视频号助手 - 直播 - 开放能力」中获得
var (
	WindowAppId                  = "xxx"
	WindowAppSecret              = "xxx"
	WindowCallbackToken          = "xxx"
	WindowCallbackEncodingAESKey = "xxx"
)

// redis 配置
var (
	RedisAddr     = "localhost:6379"
	RedisPassword = ""
	RedisDB       = 0
)
