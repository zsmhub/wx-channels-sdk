## 微信视频号 GO SDK

> “视频号小店”在2024年9月25日正式更名为“微信小店”，API调用不受影响。

- [视频号小店-官方接口文档](https://developers.weixin.qq.com/doc/channels/API/basics/getaccesstoken.html)
- [微信小店-官方接口文档](https://developers.weixin.qq.com/doc/store/API/basics/getaccesstoken.html)

Go语言实现微信视频号sdk（自建应用维度），集成了视频号小店和视频号橱窗的功能，使用简单，扩展灵活

- 支持一键生成sdk代码，包括api和回调事件
- 用缓存方案实现分布式 access_token，保证在多个服务中只有一个服务能成功调用API请求 access_token，减少API调用次数和服务重启需要重新获取的情况
    + 缓存方案支持自定义存储，默认内存存储
- 支持自定义日志存储，提供Logger interface：用于自行实现日志记录器，便于收集日志
    + 默认 log.Printf 输出

[点击查看视频号的开发博文](https://zsmhub.github.io/post/%E5%AE%9E%E6%88%98%E6%A1%88%E4%BE%8B/%E5%BE%AE%E4%BF%A1%E8%A7%86%E9%A2%91%E5%8F%B7/)

### 安装命令

```sh
go get github.com/zsmhub/wx-channels-sdk
```

### 一键生成sdk代码命令（`需手动格式化代码`）

- 生成视频号回调事件代码

    ```sh
    make callback doc=文档链接

    # example
    make callback doc=https://developers.weixin.qq.com/doc/channels/API/product/callback/ProductSpuListing.html
    ```

- 生成视频号api代码

    ```sh
    make api doc=文档链接 [prefix=生成文件名前缀]

    # example
    make api doc=https://developers.weixin.qq.com/doc/channels/API/product/get.html prefix=商品
    ```

    > tip：
    > 1. Get 方式的接口量少没做兼容，请求参数需手动整理到 Req 结构体，Post 则不用
    > 2. 部分接口文档的请求/响应 json 示例，如果缺少部分字段，需手动补上
    > 3. 部分复杂的页面需要手动整理下sdk

### sdk调用示例

**强烈建议去 ./demo 文件夹查看完整示例！**

[点击查看完整demo](https://github.com/zsmhub/wx-channels-sdk/tree/main/demo)

#### 回调事件sdk调用示例

```go
// 回调事件初始化
func (callbackRepo) InitCallbackHandler() error {
	// 视频号小店回调事件解析
	if err := channels.Sdk.NewShopCallbackHandler(ShopCallbackToken, ShopCallbackEncodingAESKey); err != nil {
		return err
	}

	// 视频号橱窗回调事件解析
	if err := channels.Sdk.NewWindowCallbackHandler(WindowCallbackToken, WindowCallbackEncodingAESKey); err != nil {
	   return err
	}

	return nil
}

// 视频号小店-解析并获取回调信息
msg, err := channels.Sdk.ShopCallback.GetCallbackMsg(c.Request())

// 视频号橱窗-解析并获取回调信息
msg, err := channels.Sdk.WindowCallback.GetCallbackMsg(c.Request())


// 视频号小店-处理回调事件完整示例
func (callbackRepo) HandleShopPostReqShouest(c echo.Context) error {
	msg, err := channels.Sdk.ShopCallback.GetCallbackMsg(c.Request())
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			_ = recover()
		}()

		switch msg.MsgType {

		case callbacks.MessageTypeEvent:
			switch msg.EventType {

			case callbacks.EventTypeProductSpuAudit:
				extras, ok := msg.Extras.(callbacks.EventProductSpuAudit)
				if !ok {
					fmt.Println("获取商品审核回调参数失败")
					return
				}
				fmt.Printf("接收到商品审核事件的原始内容：%s\n", msg.OriginalMessage)
				fmt.Printf("解析商品审核事件的结构体：%+v\n", extras)
			}

		}
	}()

	return nil
}
```

#### api sdk 调用示例

```go
// API客户端初始化
func initApiHandler() error {
	// 初始化sdk参数
	channels.Sdk.InitOptions(apis.Options{
		DcsToken: DcsTokenByRedis{},
		Logger:   Logger{},
	})

	// 视频号小店API客户端初始化
	channels.Sdk.NewShopApiClient(ShopAppId, ShopAppSecret)

	// 视频号橱窗API客户端初始化
	channels.Sdk.NewWindowApiClient(WindowAppId, WindowAppSecret)

	return nil
}

// 获取 access_token
resp, err := channels.Sdk.ShopClient.GetToken()

// 获取店铺信息
resp, err := channels.Sdk.ShopClient.ExecEcBasicsInfoGet(apis.ReqEcBasicsInfoGet{})

// error code 类型强制转换
if err != nil {
	if apiError, ok := err.(*apis.ClientError); ok {
		if apiError.Code == apis.ErrCode40013 {
			return errors.New("不合法的 AppID")
	 	}
	}
	return nil, err
}

```

### 目录结构

```sh
.
├── apis                     视频号API
│   └── api.error.go         全局错误码
├── callbacks                视频号回调事件
│   └── callback_constant.go 回调事件常量定义
├── demo                     sdk调用示例
├── generate                 一键生成脚本
│   ├── api.go               一键生成视频号API脚本
│   └── callback.go          一键生成企微回调事件脚本
├── internal                 消息加解密库
├── constant.go              全局枚举值定义
└── sdk.go                   入口文件
```

### 注意点

- 如果你发现了sdk中，没有某个回调事件或某个api，可以使用一键生成sdk代码命令生成，然后提交下pr
- 注意视频号access_token不支持互刷，获取了新的access_token，旧的access_token就会很快失效

### 推荐开源项目

- [企业微信 GO SDK](https://github.com/zsmhub/workweixin)
- [抖店开放平台 GO SDK](https://github.com/zsmhub/doudian-sdk)
- [微信视频号 GO SDK](https://github.com/zsmhub/wx-channels-sdk)
- [小红书开放平台 GO SDK](https://github.com/zsmhub/xhs-sdk)